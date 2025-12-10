-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;

-- connection string:
-- postgres://denizekindaglar@localhost:5432/gator