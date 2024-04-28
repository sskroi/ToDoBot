package main

import (
	telegramClient "ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/config"
	"ToDoBot1/pkg/events/telegram"
	"ToDoBot1/pkg/processorloop"
	"ToDoBot1/pkg/storage/sqlite"
	"log"
)

func main() {
    config := config.LoadConfig()

	// creating an object for interacting with telegram api
	tgClient := telegramClient.New(config.Telegram)
	// creating an object for interacting with sqlite3 storage
	storage, err := sqlite.New(config.SQLite)
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

