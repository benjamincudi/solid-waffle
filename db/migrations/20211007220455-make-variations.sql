-- +migrate Up
CREATE TABLE variations (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name text NOT NULL,
    UNIQUE(name)
);
CREATE TABLE product_variations (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    product_id int NOT NULL REFERENCES products(id),
    variation_ids int[] NOT NULL CHECK (variation_ids <> '{}'),
    price int NOT NULL,
    is_available bool NOT NULL,
    UNIQUE(product_id, variation_ids)
);
-- +migrate Down
