-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListFeedsWithUser :many
SELECT 
    feeds.*,
    users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;
