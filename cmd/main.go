package main

import (
	tgC "ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/events/telegram"
	processorloop "ToDoBot1/pkg/processorLoop"
	"ToDoBot1/pkg/storage/sqlite"
	"flag"
	"log"
)

func main() {
	tgBotToken, dbPath := mustFlags()

	// creating an object for interacting with the telegram api
	tgClient := tgC.New(tgBotToken)
	storage, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	err = storage.Init()
	if err != nil {
		log.Fatalf("can't init storage: %s", err.Error())
	}

	processor := telegram.New(tgClient, storage)

	mainLoop := processorloop.New(processor, 100)
	err = mainLoop.Start()
	if err != nil {
		log.Fatalf("can't start main loop %s", err.Error())
	}

}

// mustToken retrieves the value of flags (-tg-token, -db-path)
func mustFlags() (string, string) {
	tgToken := flag.String("tg-token", "", "telegram bot token") // объявляем флан для получения токена при запуске программы
	dbPath := flag.String("db-path", "", "sqlite3 db file path")

	flag.Parse()

	if *tgToken == "" {
		log.Fatal("missing telegram bot token (-tg-token flag)")
	}

	if *dbPath == "" {
		log.Fatal("missing database file path (-db-path flag)")
	}

	return *tgToken, *dbPath
}
