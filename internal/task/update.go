package task

import (
	"context"

	"github.com/eswar-7116/tgo/db"
	"github.com/eswar-7116/tgo/internal/taskdb"
)

func UpdateById(ctx context.Context, id int64, task taskdb.Task) (taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return taskdb.Task{}, err
	}
	defer database.Close()

	return queries.UpdateById(ctx, taskdb.UpdateByIdParams{
		ID:     id,
		Title:  task.Title,
		Status: task.Status,
		Tag:    task.Tag,
	})
}
