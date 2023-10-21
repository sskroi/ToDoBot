package sqlite

import (
	"ToDoBot1/pkg/e"
	"ToDoBot1/pkg/storage"
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
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
		state INT DEFAULT ` + strconv.Itoa(storage.DefState) + `,
		cur_task INT DEFAULT 0
	);`
	_, err := s.db.Exec(queryUsers)
	if err != nil {
		return e.Wrap("can't create table `users`", err)
	}

	queryTasks := `CREATE TABLE IF NOT EXISTS tasks (
		task_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT,
		description TEXT,
		create_time INTEGER,
		deadline INTEGER,
		done INTEGER NOT NULL DEFAULT 0,
		UNIQUE (user_id, title)
	);`
	if _, err := s.db.Exec(queryTasks); err != nil {
		return e.Wrap("can't create table `tasks`", err)
	}

	return nil
}

// GetState returns the state of the user or error
// if can't to get the state of user
func (s *SqliteStorage) GetState(userId int) (int, error) {
	qForGetUserState := `SELECT state FROM users WHERE user_id = ?;`

	var userState int

	err := s.db.QueryRow(qForGetUserState, userId).Scan(&userState)
	if err != nil {
		return 0, e.Wrap("can't get user's state", err)
	}

	return userState, nil
}

// Add added the task in tasks table with specified UserId
func (s *SqliteStorage) Add(userId int) error {
	err := s.checkUser(userId)
	if err != nil {
		return err
	}

	qForAddTask := `INSERT INTO tasks (user_id) VALUES (?);`

	_, err = s.db.Exec(qForAddTask, userId)
	if err != nil {
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

// checkUser checks if the user exists, if not,
// creates a user with the specified UserId.
func (s *SqliteStorage) checkUser(userId int) error {
	qForCheckUser := `SELECT user_id FROM users WHERE user_id = ?;`

	var checkUserRes int

	err := s.db.QueryRow(qForCheckUser, userId).Scan(&checkUserRes)
	if err == sql.ErrNoRows {
		qForAddUser := `INSERT INTO users (user_id) VALUES (?);`
		_, err = s.db.Exec(qForAddUser, userId)
		if err != nil {
			return e.Wrap("can't create user", err)
		}
	} else if err != nil {
		return e.Wrap("can't check user", err)
	}

	return nil
}

/*
if err != nil {
		// проверяем, что ошибку можно преобразовать в тип ошибки sqlite3, если да, проверяем,
		// является ли эта ошибка ошибкой ErrConstraintUnique, если да, возвращаем кастомный тип ошибки ErrUnique1
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrUnique1
		}
		return e.Wrap("can't add task", err)
	}
*/
