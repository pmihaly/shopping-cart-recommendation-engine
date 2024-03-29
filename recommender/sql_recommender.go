package recommender

import (
	"context"

	"github.com/google/uuid"
	"mihaly.codes/cart-recommendation-engine/database"
)

type SqlRecommender struct {
	queries database.Queries
	ctx     context.Context
}

func NewSqlRecommender(queries database.Queries, ctx context.Context) (SqlRecommender, error) {
	return SqlRecommender{
		queries,
		ctx,
	}, nil
}

func (recommender SqlRecommender) Close() {
	return
}

func (recommender SqlRecommender) GetRecommendedItems(cartId uuid.UUID) ([]string, error) {
	recommendedProductRows, err := recommender.queries.RecommendProducts(recommender.ctx, cartId)
	if err != nil {
		return nil, err
	}

	var recommendedProductIds []string
	for _, row := range recommendedProductRows {
		recommendedProductIds = append(recommendedProductIds, row.ProductID)
	}

	return recommendedProductIds, err
}

func (recommender SqlRecommender) AddOrder(cartId uuid.UUID, itemIds []string) error {
	return nil
}
