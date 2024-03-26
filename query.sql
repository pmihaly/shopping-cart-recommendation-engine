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


-- name: GetProductsByIds :many
SELECT id, name, description, category, image_url, price
FROM product
WHERE id = ANY(sqlc.arg(product_ids)::text[]);


-- name: PutCart :exec
INSERT INTO cart (id)
VALUES (sqlc.arg(cart_id))
ON CONFLICT (id) DO NOTHING;

-- name: DeleteCart :exec
DELETE FROM cart
WHERE id = sqlc.arg(cart_id);

-- name: PutCartItem :exec
INSERT INTO cart_items (cart_id, product_id)
VALUES (
  sqlc.arg(cart_id),
  sqlc.arg(product_id)
)
ON CONFLICT (cart_id, product_id) DO NOTHING;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE cart_id = sqlc.arg(cart_id)
AND product_id = sqlc.arg(product_id);

-- name: GetCartItems :many
SELECT p.id, name, description, category, image_url, price
FROM product p
JOIN
    cart_items ci ON p.id = ci.product_id
JOIN
    cart c ON c.id = ci.cart_id
WHERE
    c.id = sqlc.arg(cart_id);
