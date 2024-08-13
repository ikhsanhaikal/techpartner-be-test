-- name: ListTransactionsById :many
SELECT * FROM transactions
WHERE id = ?
ORDER BY created_at;

-- name: CreateTransaction :execresult
INSERT INTO transactions (user_id, kategori_id,
nominal, deskripsi) 
VALUES (?, ?, ?, ?);

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = ?;

