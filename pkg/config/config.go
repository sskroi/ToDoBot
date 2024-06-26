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
	SQLite   sqlite.Config   `toml:"sqlite"`
	TLS      TLSConfig       `toml:"tls"`
	Server   ServerConfig    `toml:"server"`
}

type ServerConfig struct {
	HookURL string `toml:"hookURL"`
	Port    string `toml:"port"`
}

type TLSConfig struct {
	CertificatePath string `toml:"certificatePath"`
	PrivateKeyPath  string `toml:"privateKeyPath"`
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
