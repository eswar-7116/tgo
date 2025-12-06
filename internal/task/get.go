package task

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eswar-7116/tgo/db"
	"github.com/eswar-7116/tgo/internal/taskdb"
)

func GetAllTasks(ctx context.Context) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.GetAllTasks(ctx)
}

func GetTaskById(ctx context.Context, id int64) (taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return taskdb.Task{}, err
	}
	defer database.Close()

	return queries.GetTaskById(ctx, id)
}

func GetTasksByStatus(ctx context.Context, status Status) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.GetTasksByStatus(ctx, sql.NullString{String: string(status), Valid: true})
}

func GetTasksByTag(ctx context.Context, tag string) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.GetTasksByTag(ctx, sql.NullString{String: tag, Valid: tag != ""})
}

func GetTasksByTitle(ctx context.Context, title string) ([]taskdb.Task, error) {
	if title == "" {
		return nil, fmt.Errorf("No title provided to search")
	}
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.GetTasksByTitle(ctx, sql.NullString{String: title, Valid: false})
}
