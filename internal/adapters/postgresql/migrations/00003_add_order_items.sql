-- +goose Up
CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipping', 'received', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS promotions (
    id BIGSERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    discount_percentage INTEGER NOT NULL CHECK (discount_percentage >= 0),
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_deleted BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    promotion_id BIGINT REFERENCES promotions(id),
    base_price_in_cents INTEGER NOT NULL CHECK (base_price_in_cents >= 0),
    total_discount_in_cents INTEGER NOT NULL CHECK (total_discount_in_cents >= 0),
    total_price_in_cents INTEGER NOT NULL CHECK (total_price_in_cents >= 0),
    status order_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE products RENAME COLUMN quantity TO stock;

CREATE TABLE IF NOT EXISTS order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL,
    total_price_per_checkout_in_cents INTEGER NOT NULL CHECK (total_price_per_checkout_in_cents >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
ALTER TABLE products RENAME COLUMN stock TO quantity;
DROP TABLE carts;
DROP TABLE promotions;
DROP TABLE order_items;
DROP TABLE orders;
DROP TYPE order_status;
