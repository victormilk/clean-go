-- name: GetTodo :one
SELECT *
FROM todos
WHERE id = $1 LIMIT 1;

-- name: GetTodos :many
SELECT *
FROM todos
ORDER BY title ASC;

-- name: CreateTodo :exec
INSERT INTO todos (id, title, description, completed, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateTodo :exec
UPDATE todos
SET title       = $2,
    description = $3,
    completed   = $4,
    updated_at  = $5
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE
FROM todos
WHERE id = $1;

-- name: ExistsTodo :one
SELECT EXISTS(SELECT 1 FROM todos WHERE id = $1 LIMIT 1);