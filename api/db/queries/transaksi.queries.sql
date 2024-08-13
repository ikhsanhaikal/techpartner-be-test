-- name: ListTransactionsByUserIdDefault :many
SELECT * FROM transactions 
WHERE user_id = ? AND YEAR(created_at) = YEAR(CURRENT_DATE()) AND 
      MONTH(created_at) = MONTH(CURRENT_DATE());

-- name: GetTransactions :one
SELECT * FROM transactions
WHERE id = ?;

-- name: ListTransactionsByUserIdRange :many
SELECT * FROM transactions 
WHERE created_at BETWEEN ? AND ? AND user_id = ?; 

-- name: CreateTransaction :execresult
INSERT INTO transactions (user_id, kategori_id,
nominal, deskripsi) 
VALUES (?, ?, ?, ?);

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = ?;

