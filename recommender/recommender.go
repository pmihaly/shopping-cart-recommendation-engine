package recommender

import "github.com/google/uuid"

type Recommender interface {
	GetRecommendedItems(cartId uuid.UUID) ([]string, error)
	AddOrder(cartId uuid.UUID, itemIds []string) error
	Close()
}
