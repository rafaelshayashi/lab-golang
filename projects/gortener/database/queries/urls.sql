-- name: GetURL :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;

-- name: CreateURL :one
INSERT INTO urls (
  url, short_url, hash, user_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id;