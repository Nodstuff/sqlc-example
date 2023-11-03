-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY name;

-- name: CreateBook :one
INSERT INTO books (
    name, bio
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: UpdateBook :one
UPDATE books
set name = $2,
    bio = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;