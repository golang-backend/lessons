-- +goose Up
CREATE TABLE IF NOT EXISTS tasks(
    id          SERIAL PRIMARY KEY,
    user_id     INT REFERENCES users(id),
    title       VARCHAR(255),
    description TEXT,
    status      VARCHAR(55), -- new, processing, done, failed
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF NOT EXISTS tasks;
