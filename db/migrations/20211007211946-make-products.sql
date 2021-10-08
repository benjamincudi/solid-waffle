-- +migrate Up
CREATE TABLE products (
     id SERIAL PRIMARY KEY,
     created_at timestamp without time zone NOT NULL,
     updated_at timestamp without time zone NOT NULL,
     seller_id int NOT NULL REFERENCES sellers(id),
     name text NOT NULL,
     UNIQUE(seller_id, name)
);
-- +migrate Down
