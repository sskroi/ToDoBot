package config

import (
	"github.com/BurntSushi/toml"

	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/storage/postgres"
	"ToDoBot1/pkg/storage/sqlite"
)

const (
    configPath = "./configs/config.toml"
)

type Config struct {
    Postgres postgres.Config `toml:"postgres"`
    Telegram telegram.Config `toml:"telegram"`
    SQLite sqlite.Config `toml:"sqlite"`
}

func LoadConfig() *Config {
    cfg := new(Config)

    _, err := toml.DecodeFile(configPath, cfg)
    if err != nil {
        panic(err)
    }

    return cfg
}


