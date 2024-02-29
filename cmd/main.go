package main

import (
	tgC "ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/events/telegram"
	"ToDoBot1/pkg/processorloop"
	"ToDoBot1/pkg/storage/sqlite"
	"fmt"
	"log"
	"os"
)

func main() {
	tgBotToken, dbPath := loadCfgFromEnv()

	// creating an object for interacting with telegram api
	tgClient := tgC.New(tgBotToken)
	// creating an object for interacting with sqlite3 storage
	storage, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	processor := telegram.New(tgClient, storage)

	mainLoop := processorloop.New(processor, 100)
	err = mainLoop.Start()
	if err != nil {
		log.Fatalf("can't start main loop %s", err.Error())
	}

}

// mustFlags retrieves the value of flags (-tg-token, -db-path)
func loadCfgFromEnv() (string, string) {
	tgToken := os.Getenv("TG_TOKEN")
	if tgToken == "" {
		log.Fatal("missing telegram bot token 'TG_TOKEN' env var")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Print("database path (DB_PATH env) not set, using default value = ./database/db1.sqlite3")

		if err := initDefaultDatabaseFile(); err != nil {
			log.Fatal(err.Error())
		}

		dbPath = "./database/db1.sqlite3"
	}

	return tgToken, dbPath
}

func initDefaultDatabaseFile() error {
	if _, err := os.Stat("./database"); os.IsNotExist(err) {
		err = os.Mkdir("./database", 0755)
		if err != nil {
			return fmt.Errorf("can't create './database' folder: %w", err)
		}
	}

	if _, err := os.Stat("./database/db1.sqlite3"); os.IsNotExist(err) {
		file, err := os.Create("./database/db1.sqlite3")
		if err != nil {
			return fmt.Errorf("can't create database file './datavase/db1.sqlite3': %w", err)
		}
		file.Close()
	}
	return nil
}
