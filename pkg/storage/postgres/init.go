package postgres

import "ToDoBot1/pkg/e"

func (self *PostgresDB) Init() error {
	initUsersQuery := `CREATE TABLE IF NOT EXISTS users (
		user_id BIGINT PRIMARY KEY,
		username VARCHAR(64),
		state INT DEFAULT 0,
		cur_task BIGINT,
		timezone VARCHAR(128)
	);`

	_, err := self.db.Exec(initUsersQuery)
	if err != nil {
		return e.Wrap("Postgres init: ", err)
	}

	initTasksQuery := `CREATE TABLE IF NOT EXISTS tasks (
		task_id BIGSERIAL,
		user_id BIGINT NOT NULL,
		name VARCHAR(255),
		create_time BIGINT,
		deadline BIGINT,
		done SMALLINT NOT NULL DEFAULT 0,
		finish_time BIGINT,
		UNIQUE(user_id, name)
	);`
	_, err = self.db.Exec(initTasksQuery)
	if err != nil {
		return e.Wrap("Postgres init: ", err)
	}

	return nil
}
