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
SELECT * FROM users WHERE name = $1;

-- name: Reset :exec
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

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetFeedCreator :one
SELECT name FROM users where id = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
    ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
    ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedIDFromUrl :one
SELECT id FROM feeds WHERE url = $1;

-- name: GetFeedFromUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedFollowForUser :many
SELECT
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
INNER JOIN users
    ON feed_follows.user_id = users.id
INNER JOIN feeds
    ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
WITH feed_follow_from_url AS (
    SELECT id FROM feeds WHERE url = $1
)
DELETE FROM feed_follows 
WHERE feed_id = (SELECT id FROM feed_follow_from_url);

-- name: MarkFeedFetched :exec
UPDATE feeds
SET 
    last_fetched_at = $1,
    updated_at = $2
WHERE 
    id = $3;

-- name: GetNextFeedToFetch :one
SELECT *
    FROM feeds 
    ORDER BY last_fetched_at ASC NULLS FIRST;

-- name: CreatePost :exec
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);

-- name: GetPostsForUser :many
SELECT * FROM posts
ORDER BY published_at DESC
LIMIT $1;

-- name: UrlExists :one
SELECT 1 
FROM posts
WHERE url = $1
LIMIT 1;