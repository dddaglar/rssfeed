-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4  
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name
FROM users
WHERE name = $1;

-- name: ResetTable :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT name FROM users;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds;

-- name: GetUserByID :one
SELECT id, created_at, updated_at, name
FROM users
WHERE id = $1;

-- name: CreateFeedFollow :one
WITH iff AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
) SELECT
    iff.*, feeds.name AS feed_name, users.name AS user_name
    FROM iff
    INNER JOIN feeds ON iff.feed_id = feeds.id
    INNER JOIN users ON iff.user_id = users.id;

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
WITH ff AS (
    SELECT id, created_at, updated_at, user_id, feed_id 
    FROM feed_follows
    WHERE feed_follows.user_id = $1
) SELECT 
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id AS follow_user_id,
    ff.feed_id,
    feeds.name AS feed_name, 
    users.name AS user_name
    FROM ff
    INNER JOIN feeds ON ff.feed_id = feeds.id
    INNER JOIN users ON ff.user_id = users.id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 )
RETURNING *;

-- name: GetPostsForUser :many
SELECT * 
FROM posts
WHERE feed_id IN (
    SELECT feed_id 
    FROM feed_follows 
    WHERE user_id = $1
)
ORDER BY published_at DESC NULLS LAST;