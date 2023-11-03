-- name: GetPublisher :one
SELECT * FROM publishers
WHERE id = $1 LIMIT 1;

-- name: ListPublishers :many
SELECT * FROM publishers
ORDER BY name;

-- name: CreatePublisher :one
INSERT INTO publishers (
    name, bio
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: UpdatePublisher :one
UPDATE publishers
set name = $2,
    bio = $3
WHERE id = $1
    RETURNING *;

-- name: DeletePublisher :exec
DELETE FROM publishers
WHERE id = $1;