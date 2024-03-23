// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package database

import (
	"context"
)

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, category, image_url, price FROM product
LIMIT 100
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.ImageUrl,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
