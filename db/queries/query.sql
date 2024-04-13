-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id;

-- name: UpdateAccount :exec
UPDATE accounts
  set balance = $2
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListEntriesByAccount :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: UpdateEntry :exec
UPDATE entries
  set amount = $2
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListTransfersByFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id;

-- name: ListTransfersByToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id;

-- name: ListTransfersByFromAndToAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1 and to_account_id = $2
ORDER BY id;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: UpdateTransfer :exec
UPDATE transfers
  set amount = $2
WHERE id = $1;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;