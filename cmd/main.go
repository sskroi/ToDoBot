package main

import (
	"ToDoBot1/pkg/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// создание объекта для взаимодействия с api телеграма
	tgClient := telegram.New(tgBotHost, mustToken())
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
