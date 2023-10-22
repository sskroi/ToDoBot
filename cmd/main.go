package main

import (
	tgClient "ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/events/telegram"
	processorloop "ToDoBot1/pkg/processorLoop"
	"ToDoBot1/pkg/storage/sqlite"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "./db/db1.sqlite3"
)

func main() {
	// создание объекта для взаимодействия с api телеграма
	tgClient := tgClient.New(tgBotHost, mustToken())
	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	err = storage.Init()
	if err != nil {
		log.Fatalf("can't init storage: %s", err.Error())
	}

	processor := telegram.New(&tgClient, storage)

	mainLoop := processorloop.New(processor, 100)
	err = mainLoop.Start()
	if err != nil {
		log.Fatalf("can't start main loop %s", err.Error())
	}

}

// mustToken извлекает значение токена из флага tg-token
func mustToken() string {
	token := flag.String("tg-token", "", "telegram bot token") // объявляем флан для получения токена при запуске программы

	flag.Parse()

	if *token == "" {
		log.Fatal("Отсутствует токен для бота")
	}

	return *token
}
