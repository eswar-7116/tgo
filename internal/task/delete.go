package task

import (
	"context"
	"database/sql"

	"github.com/eswar-7116/tgo/db"
	"github.com/eswar-7116/tgo/internal/taskdb"
)

func DeleteAllTasks(ctx context.Context) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.DeleteAllTasks(ctx)
}

func DeleteById(ctx context.Context, id int64) (taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return taskdb.Task{}, err
	}
	defer database.Close()

	return queries.DeleteById(ctx, id)
}

func DeleteByStatus(ctx context.Context, status Status) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.DeleteByStatus(ctx, sql.NullString{String: string(status), Valid: true})
}

func DeleteByTag(ctx context.Context, tag string) ([]taskdb.Task, error) {
	database, queries, err := db.OpenDB(ctx)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	return queries.DeleteByTag(ctx, sql.NullString{String: tag, Valid: tag != ""})
}
