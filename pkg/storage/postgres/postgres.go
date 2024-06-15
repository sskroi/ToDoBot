package postgres

import (
	"ToDoBot1/pkg/storage"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

type PostgresDB struct {
	db *sqlx.DB
}

func New(cfg Config) (*PostgresDB, error) {
	const fn = "postgres.New"

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &PostgresDB{db: db}, nil
}

func (d *PostgresDB) GetState(userId uint64) (int, error) {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) SetState(userId uint64, state int) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) Add(userId uint64) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) UpdTitle(userId uint64, title string) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) UpdDescription(userId uint64, description string) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) UpdDeadline(userId uint64, deadline uint64, createTime uint64) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) Delete(userId uint64, title string) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) CloseTask(userId uint64, title string, finishTime uint64) error {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) Uncompl(userId uint64) ([]storage.Task, error) {
	panic("not implemented") // TODO: Implement
}
func (d *PostgresDB) Compl(userId uint64) ([]storage.Task, error) {
	panic("not implemented") // TODO: Implement
}
