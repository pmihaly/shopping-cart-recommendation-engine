package main

import (
	"context"
	"log/slog"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ExtraItemRecommender struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func NewExtraItemRecommender(username, password, uri string, ctx context.Context) (ExtraItemRecommender, error) {
	auth := neo4j.BasicAuth(username, password, "")
	driver, err := neo4j.NewDriverWithContext(uri, auth)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		slog.Error("failed to connect to neo4j", err, "uri", uri)
		return ExtraItemRecommender{}, err
	}

	return ExtraItemRecommender{
		driver,
		ctx,
	}, nil
}

func (recommender ExtraItemRecommender) Close() {
	recommender.driver.Close(recommender.ctx)
}

func (recommender ExtraItemRecommender) GetRecommendedExtraItems(cartId string) ([]string, error) {
	result, err := neo4j.ExecuteQuery(recommender.ctx, recommender.driver,
		`
		MATCH (s1:Cart)-[:ORDERED]->(p:Product)
		WHERE s1.cartId = $cartId
		WITH s1, collect(DISTINCT p) AS products1

		MATCH (s2:Cart)-[:ORDERED]->(p:Product)
		WHERE s1 <> s2
		WITH s1, s2, products1, collect(DISTINCT p) AS products2

		WITH s1, s2, products1, products2, apoc.coll.intersection(products1, products2) AS intersection, apoc.coll.union(products1, products2) AS union

		WITH s1, s2, intersection, union, size(intersection) * 1.0 / size(union) AS jaccard_index

		ORDER BY jaccard_index DESC, s2.cartId
		WITH s1, collect(s2)[..$neighborsCount] AS neighbors
		WHERE size(neighbors) = $neighborsCount

		UNWIND neighbors AS neighbor

		MATCH (neighbor)-[:ORDERED]->(p:Product)
		WHERE NOT (s1)-[:ORDERED]->(p)

		WITH s1, p, count(DISTINCT neighbor) AS countnns
		ORDER BY s1.cartId, countnns DESC

		RETURN collect(p.productId)[..$recommendationCount] AS recommendations
		`,
		map[string]any{
			"cartId":              cartId,
			"neighborsCount":      25,
			"recommendationCount": 5,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		slog.Error("failed to run neo4j query", err, "cartId", cartId)
		return make([]string, 0), err
	}

	recommendedProductIds := []string{}
	for _, record := range result.Records {
		recommendations, _ := record.Get("recommendations")
		if recs, ok := recommendations.([]interface{}); ok {
			for _, rec := range recs {
				if str, ok := rec.(string); ok {
					recommendedProductIds = append(recommendedProductIds, str)
				}
			}
		}
	}

	return recommendedProductIds, nil
}

func (recommender ExtraItemRecommender) AddOrder(cartId string, itemIds []string) error {
	_, err := neo4j.ExecuteQuery(recommender.ctx, recommender.driver,
		`
		MERGE (cart:Cart {cartId: $cartId})
		WITH cart
		UNWIND $itemIds AS itemId
		MERGE (product:Product {productId: itemId})
		MERGE (cart)-[:ORDERED]->(product)
		`,
		map[string]interface{}{
			"cartId":  cartId,
			"itemIds": itemIds,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)

	if err != nil {
		slog.Error("failed to add order to neo4j", err, "cartId", cartId)
		return err
	}

	return nil
}
