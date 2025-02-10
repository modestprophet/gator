-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (url) DO NOTHING
RETURNING *;

-- name: GetPostsForUser :many
SELECT 
    posts.id,
    posts.title,
    posts.url,
    posts.description,
    posts.published_at,
    posts.feed_id,
    feeds.url AS feed_url
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON posts.feed_id = feeds.id
WHERE users.name = $1
ORDER BY posts.published_at DESC
LIMIT $2;