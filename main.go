package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"mihaly.codes/shopping-cart-recommendation-engine/database"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	ctx1 := context.Background()
	auth := neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), "")
	driver, err := neo4j.NewDriverWithContext(os.Getenv("NEO4J_URI"), auth)
	defer driver.Close(ctx1)

	err = driver.VerifyConnectivity(ctx1)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, req *http.Request) {
		products, err := queries.ListProducts(ctx, database.ListProductsParams{Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/search", func(w http.ResponseWriter, req *http.Request) {
		searchValue := req.URL.Query().Get("q")

		products, err := queries.SearchProducts(ctx, database.SearchProductsParams{SearchValue: searchValue, Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/recommended", func(w http.ResponseWriter, req *http.Request) {
		shoppingCartId := req.URL.Query().Get("shoppingCartId")

		result, err := neo4j.ExecuteQuery(ctx, driver,
			`
			MATCH (s1:ShoppingCart)-[:ORDERED]->(p:Product)
			WHERE s1.shoppingCartId = $shoppingCartId
			WITH s1, collect(DISTINCT p) AS products1

			MATCH (s2:ShoppingCart)-[:ORDERED]->(p:Product)
			WHERE s1 <> s2
			WITH s1, s2, products1, collect(DISTINCT p) AS products2

			WITH s1, s2, products1, products2, apoc.coll.intersection(products1, products2) AS intersection, apoc.coll.union(products1, products2) AS union

			WITH s1, s2, intersection, union, size(intersection) * 1.0 / size(union) AS jaccard_index

			ORDER BY jaccard_index DESC, s2.shoppingCartId
			WITH s1, collect(s2)[..$k] AS neighbors
			WHERE size(neighbors) = $k

			UNWIND neighbors AS neighbor

			MATCH (neighbor)-[:ORDERED]->(p:Product)
			WHERE NOT (s1)-[:ORDERED]->(p)

			WITH s1, p, count(DISTINCT neighbor) AS countnns
			ORDER BY s1.shoppingCartId, countnns DESC

			RETURN collect(p.productId)[..$n] AS recommendations
			`,
			map[string]any{
				"shoppingCartId": shoppingCartId,
				"k":              25,
				"n":              5,
			}, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))

		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
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

		products, err := queries.GetProductsByIds(ctx, recommendedProductIds)
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	port := ":8090"
	log.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
