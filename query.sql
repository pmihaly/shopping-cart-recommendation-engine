-- name: ListProducts :many
SELECT id, name, description, category, image_url, price
FROM product
ORDER BY RANDOM()
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



-- name: RecommendProducts :many
WITH UserSimilarity AS (
  SELECT
    ci1.cart_id AS user1_cart_id,
    ci2.cart_id AS user2_cart_id,
    COUNT(*) AS common_items_count,
    (SELECT COUNT(*) FROM cart_items WHERE cart_id = ci1.cart_id) AS user1_total_items,
    (SELECT COUNT(*) FROM cart_items WHERE cart_id = ci2.cart_id) AS user2_total_items
  FROM
    cart_items ci1
  JOIN
    cart_items ci2 ON ci1.product_id = ci2.product_id AND ci1.cart_id <> ci2.cart_id
  WHERE
    ci1.cart_id = sqlc.arg(cart_id)
  GROUP BY
    ci1.cart_id, ci2.cart_id
),

UserCartSimilarity AS (
  SELECT
    user1_cart_id,
    user2_cart_id,
    common_items_count::float / (user1_total_items + user2_total_items - common_items_count) AS similarity
  FROM
    UserSimilarity
),

Recommendations AS (
  SELECT
    ci.product_id,
    SUM(UserCartSimilarity.similarity) AS total_similarity
  FROM
    cart_items ci
  JOIN
    UserCartSimilarity ON ci.cart_id = UserCartSimilarity.user2_cart_id
  WHERE
    ci.product_id NOT IN (SELECT product_id FROM cart_items WHERE cart_id = sqlc.arg(cart_id))
  GROUP BY
    ci.product_id
)

SELECT
  r.product_id,
  r.total_similarity::int
FROM
  Recommendations r
ORDER BY
  r.total_similarity DESC
LIMIT
  10;
  -- sqlc.arg(recommendation_limit);
