package sqlite3

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	db *sql.DB
}

func New(path string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	return &SqliteStorage{
		db: db,
	}, nil
}

func (s *SqliteStorage) Init() error {
	queryUsers := `CREATE TABLE IF NOT EXISTS users (
		user_id INT PRIMARY KEY,
		username VARCHAR(255)
	);`
	_, err := s.db.Exec(queryUsers)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	queryTasks := `CREATE TABLE IF NOT EXISTS tasks (
		task_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT,
		create_time INTEGER NOT NULL,
		deadline INTEGER NOT NULL,
		done INTEGER NOT NULL DEFAULT 0,
		UNIQUE (user_id, title)
	);`
	if _, err := s.db.Exec(queryTasks); err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
