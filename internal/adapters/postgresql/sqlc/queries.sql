-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (id, name, price_in_cents, quantity, created_at)
VALUES (nextval('products_id_seq'), $1, $2, $3, now())
RETURNING id;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price_in_cents = $2, quantity = $3
WHERE id = $4;