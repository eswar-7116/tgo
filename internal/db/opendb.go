package db

import (
	"context"
	"database/sql"
	_ "embed"
	"os"
	"path"

	"github.com/eswar-7116/tgo/internal/taskdb"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func OpenDB(ctx context.Context) (*sql.DB, *taskdb.Queries, error) {
	// Determine the OS-specific cache directory
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, nil, err
	}
	// Path to tgo's cache folder
	cachePath := path.Join(cacheDir, "tgo")

	// Create cache directory if not exists
	if err := os.MkdirAll(cachePath, os.ModePerm); err != nil {
		return nil, nil, err
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite", path.Join(cachePath, "tasks.db"))
	if err != nil {
		return nil, nil, err
	}

	// Execute the DDL
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, nil, err
	}

	// Initialize sqlc-generated query helpers
	queries := taskdb.New(db)
	return db, queries, nil
}
