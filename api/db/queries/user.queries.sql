-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: CreateUser :execresult
INSERT INTO users (name, email, password) 
VALUES (?, ?, ?);