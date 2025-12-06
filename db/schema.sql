CREATE TABLE IF NOT EXISTS
    tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        tag TEXT,
        status VARCHAR(5) DEFAULT 'TODO',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CHECK (status IN ('TODO', 'DOING', 'DONE'))
    );

CREATE TRIGGER IF NOT EXISTS update_tasks_updated_at AFTER
UPDATE ON tasks FOR EACH ROW BEGIN
UPDATE tasks
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;
