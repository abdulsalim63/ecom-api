-- CRUD Products
-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (id, name, price_in_cents, stock, created_at)
VALUES (nextval('products_id_seq'), $1, $2, $3, now())
RETURNING id;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price_in_cents = $2, stock = $3
WHERE id = $4;


-- CRUD Users
-- name: CreateUser :one
INSERT INTO users (id, name, username, created_at)
VALUES (nextval('users_id_seq'), $1, $2, now())
RETURNING id;

-- name: ListUsers :many
SELECT * FROM users;

-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = $1
WHERE id = $2;


-- CRUD Promotions
-- name: ListPromotions :many
SELECT * FROM promotions
WHERE is_deleted = false;

-- name: FindPromotionByID :one
SELECT * FROM promotions
WHERE id = $1;

-- name: CreatePromotion :one
INSERT INTO promotions (id, code, description, discount_percentage, start_date, end_date, created_at)
VALUES (nextval('promotions_id_seq'), $1, $2, $3, $4, $5, now())
RETURNING id;

-- name: DeletePromotion :exec
UPDATE promotions
SET is_deleted = true
WHERE id = $1;

-- name: UpdatePromotion :exec
UPDATE promotions
SET code = $1, description = $2, discount_percentage = $3, start_date = $4, end_date = $5
WHERE id = $6;


-- CRUD Carts
-- name: AddToCart :one
INSERT INTO carts (id, user_id, product_id, quantity, created_at)
VALUES (nextval('carts_id_seq'), $1, $2, $3, now())
RETURNING id;

-- name: UpdateCartItem :exec
UPDATE carts
SET quantity = $1, updated_at = now()
WHERE id = $2;

-- name: RemoveCartItem :exec
UPDATE carts
SET is_deleted = true, updated_at = now()
WHERE id = $1;

-- name: ListCartItems :many
SELECT * FROM carts
WHERE user_id = $1 AND is_deleted = false
ORDER BY COALESCE(updated_at, created_at) DESC
LIMIT 10 OFFSET $2;


-- Checkout & CRUD Orders
-- name: Checkout :exec
-- Create order data, migrate checked_out_items to order_items table
-- Update stock on products table and soft delete item in carts table
WITH checked_out_items AS (
  SELECT carts.product_id, carts.quantity, products.price_in_cents * carts.quantity AS based_price
  FROM carts
  INNER JOIN products ON carts.product_id = products.id
  WHERE carts.user_id = $1 AND carts.id IN ($2)
),
inserted_orders AS (
  INSERT INTO orders (id, user_id, base_price_in_cents, total_discount_in_cents, total_price_in_cents, status, created_at)
  SELECT nextval('orders_id_seq'), $1, 
    SUM(based_price), 
    $2 * SUM(based_price) / 100, 
    SUM(based_price) - $2 * SUM(based_price) / 100, 
    'pending', now()
  FROM checked_out_items
),
inserted_order_items AS (
  INSERT INTO order_items (id, order_id, product_id, quantity, total_price_per_checkout_in_cents, created_at)
  SELECT nextval('order_items_id_seq'), inserted_orders.id, checked_out_items.product_id, checked_out_items.quantity, checked_out_items.based_price * $2 / 100, now()
  FROM checked_out_items, inserted_orders
),
updated_carts AS (
  UPDATE carts
  SET is_deleted = true, updated_at = now()
  WHERE id IN ($2)
)
UPDATE products
SET stock = stock - checked_out_items.quantity
FROM checked_out_items
WHERE products.id = checked_out_items.product_id;

-- name: UpdateOrder :exec
UPDATE orders
SET status = $1, updated_at = now()
WHERE id = $2;

-- name: ListOrders :many
SELECT * FROM orders
WHERE user_id = $1 AND ($2 IS NULL OR $2 = '' OR status = $2)
ORDER BY COALESCE(updated_at, created_at) DESC;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1;

-- name: ListOrderItems :many
SELECT * FROM order_items
WHERE order_id = $1;

-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE id = $1;