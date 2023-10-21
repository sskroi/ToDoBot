package sqlite

import (
	"ToDoBot1/pkg/e"
	"ToDoBot1/pkg/storage"
	"database/sql"
	"errors"
	"strconv"

	"github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	db *sql.DB
}

var (
	ErrUnique1      = errors.New("unique error")
	ErrNotExist     = errors.New("requested data does not exist")
	ErrAlreayClosed = errors.New("task alreay closed")
)

// New устанавливает соединение с файлом БД и возвращает
// объект для взимодействия с базой данных sqlite3.
// Возвращает ошибку, если не удалось открыть файл с БД.
func New(path string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, e.Wrap("can't open database", err)
	}

	if err := db.Ping(); err != nil {
		return nil, e.Wrap("can't open database", err)
	}

	return &SqliteStorage{
		db: db,
	}, nil
}

// Init инициализирует базу данных
// (создёт таблицы,если они не были созданы)
func (s *SqliteStorage) Init() error {
	queryUsers := `CREATE TABLE IF NOT EXISTS users (
		user_id INT PRIMARY KEY,
		username VARCHAR(255),
		state INT DEFAULT ` + strconv.Itoa(storage.DefState) + `
	);`
	_, err := s.db.Exec(queryUsers)
	if err != nil {
		return e.Wrap("can't create table `users`", err)
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
		return e.Wrap("can't create table `tasks`", err)
	}

	return nil
}

// Add добавляет задачу в таблицу tasks. Если пользователя, который добавляет
// задачу, нету в таблице users - добавляет его в таблицу users.
func (s *SqliteStorage) Add(task *storage.Task) error {
	qForCheckUser := `SELECT user_id FROM users WHERE user_id = ?;`

	var CheckUserRes int

	err := s.db.QueryRow(qForCheckUser, task.UserId).Scan(&CheckUserRes)
	if err == sql.ErrNoRows {
		qForAddUser := `INSERT INTO users (user_id) VALUES (?);`
		_, err = s.db.Exec(qForAddUser, task.UserId)
		if err != nil {
			return e.Wrap("can't create user", err)
		}
	} else if err != nil {
		return e.Wrap("can't check user", err)
	}

	qForAddTask := `INSERT INTO tasks (user_id, title, description, create_time, deadline)
					VALUES (?, ?, ?, ?, ?)`
	_, err = s.db.Exec(qForAddTask, task.UserId, task.Title, task.Description, task.CreateTime, task.Deadline)
	if err != nil {
		/* проверяем, что ошибку можно преобразовать в тип ошибки sqlite3, если да, проверяем,
		является ли эта ошибка ошибкой ErrConstraintUnique, если да, возвращаем кастомный тип ошибки ErrUnique1*/
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrUnique1
		}
		return e.Wrap("can't add task", err)
	}

	return nil
}

func (s *SqliteStorage) Delete(task *storage.Task) error {
	err := s.isTaskExist(task)
	if err == ErrNotExist {
		return err
	} else if err != nil {
		return err
	}

	qForDelTask := `DELETE FROM tasks WHERE user_id = ? AND title = ?;`
	_, err = s.db.Exec(qForDelTask, task.UserId, task.Title)

	if err != nil {
		return e.Wrap("can't delete task", err)
	}

	return nil
}

func (s *SqliteStorage) isTaskExist(task *storage.Task) error {
	qForCheckExist := `SELECT task_id FROM tasks WHERE user_id = ? AND title = ?;`

	var checkExistRes int

	err := s.db.QueryRow(qForCheckExist, task.UserId, task.Title).Scan(&checkExistRes)
	if err == sql.ErrNoRows {
		return ErrNotExist
	} else if err != nil {
		return e.Wrap("can't delete task", err)
	}

	return nil
}
