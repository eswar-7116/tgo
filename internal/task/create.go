package task

import (
	"context"
	"database/sql"

	"github.com/eswar-7116/tgo/db"
	"github.com/eswar-7116/tgo/internal/taskdb"
)

func CreateTask(ctx context.Context, title string, tag string) (taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return taskdb.Task{}, err
	}
	defer database.Close()

	return queries.CreateTask(ctx, taskdb.CreateTaskParams{
		Title: title,
		Tag:   sql.NullString{String: tag, Valid: tag != ""},
	})
}
