-- name: ListCategories :many
SELECT * FROM categories
ORDER BY tipe;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = ?;

-- name: CreateCategory :execresult
INSERT INTO categories (nama, tipe, deskripsi) 
VALUES (?, ?, ?);

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;


