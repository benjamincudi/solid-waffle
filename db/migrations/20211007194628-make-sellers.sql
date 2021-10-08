-- +migrate Up
CREATE TABLE sellers (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name text NOT NULL,
    UNIQUE(name)
);
-- +migrate Down
