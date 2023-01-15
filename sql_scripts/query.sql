-- name: GetUrlFromShortId :one
SELECT * FROM urls
WHERE short_id = $1
LIMIT 1;

-- name: CreateShortUrl :one
INSERT INTO urls (
    short_id,short_url,long_url
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUrlCount :exec
UPDATE urls 
    SET redirect_count = $2
WHERE short_id = $1;