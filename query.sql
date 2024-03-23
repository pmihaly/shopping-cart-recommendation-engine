-- name: ListProducts :many
SELECT id, name, description, category, image_url, price
FROM product
OFFSET sqlc.arg(skip)
LIMIT sqlc.arg(take);

-- name: SearchProducts :many
SELECT id, name, description, category, image_url, price
FROM product
WHERE name_search @@ websearch_to_tsquery('simple', sqlc.arg(search_value))
   OR description_search @@ websearch_to_tsquery('simple', sqlc.arg(search_value))
OFFSET sqlc.arg(skip)
LIMIT sqlc.arg(take);
