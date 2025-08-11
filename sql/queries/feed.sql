-- name: CreateFeed :one
INSERT INTO feed (
  id,
  created_at,
  updated_at,
  url,
  name,
  user_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT feed.*, users.name AS username FROM feed JOIN users ON user_id = users.id;

-- name: GetFeedWithUrl :one
SELECT * FROM feed WHERE url = $1;

-- name: FollowFeed :one
INSERT INTO feed_follows (
  id,
  created_at,
  updated_at,
  user_id,
  feed_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feed.name AS feed_name
FROM feed_follows
JOIN feed ON feed_id = feed.id
WHERE feed_follows.user_id = $1;
