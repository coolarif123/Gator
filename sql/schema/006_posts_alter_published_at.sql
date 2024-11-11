-- +goose Up
ALTER TABLE posts
ALTER COLUMN published_at TYPE TEXT;

-- +goose Down
ALTER TABLE posts
ALTER COLUMN published_at TYPE TIMESTAMP;