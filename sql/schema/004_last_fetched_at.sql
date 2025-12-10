-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched_at;

-- connection string:
-- postgres://denizekindaglar@localhost:5432/gator