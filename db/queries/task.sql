-- name: CreateTask :one
INSERT INTO
    tasks (title, tag)
VALUES
    (?, ?) RETURNING *;

-- name: GetAllTasks :many
SELECT
    *
FROM
    tasks
ORDER BY
    CASE
        WHEN ? = 'updated' THEN updated_at
        ELSE created_at
    END DESC;

-- name: GetTaskById :one
SELECT
    *
FROM
    tasks
WHERE
    id = ?;

-- name: GetTasksByTag :many
SELECT
    *
FROM
    tasks
WHERE
    tag = ?
ORDER BY
    CASE
        WHEN ? = 'updated' THEN updated_at
        ELSE created_at
    END DESC;

-- name: GetTasksByStatus :many
SELECT
    *
FROM
    tasks
WHERE
    status = ?
ORDER BY
    CASE
        WHEN ? = 'updated' THEN updated_at
        ELSE created_at
    END DESC;

-- name: DeleteAllTasks :many
DELETE FROM tasks RETURNING *;

-- name: DeleteById :one
DELETE FROM tasks
WHERE
    id = ? RETURNING *;

-- name: DeleteByStatus :many
DELETE FROM tasks
WHERE
    status = ? RETURNING *;

-- name: DeleteByTag :many
DELETE FROM tasks
WHERE
    status = ? RETURNING *;

-- name: UpdateById :one
UPDATE tasks
SET
    title = ?,
    status = ?,
    tag = ?
WHERE
    id = ? RETURNING *;

-- name: GetTasksByTitle :many
SELECT
    *
FROM
    tasks
WHERE
    title LIKE '%' || ? || '%'
ORDER BY
    CASE
        WHEN ? = 'updated' THEN t.updated_at
        ELSE t.created_at
    END DESC;
