-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, town, postal_code, created_by
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetProperty :one
SELECT * FROM properties
JOIN users ON properties.created_by = users.id
WHERE properties.id = $1
LIMIT 1;

-- name: FindByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;
