-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(255),
    last_name   VARCHAR(255),
    phone       VARCHAR(20),
    balance     DECIMAL(10, 2) DEFAULT 0.00,
    created_at  TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF NOT EXISTS users;
