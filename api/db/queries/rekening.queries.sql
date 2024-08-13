-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = ? AND user_id = ?;

-- name: Deposit :exec
UPDATE accounts
SET saldo = saldo + ?
WHERE id = ?;

-- name: Withdraw :exec
UPDATE accounts
SET saldo = saldo - ?
WHERE id = ?;