package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	token := mustToken()
	fmt.Println(token)
}

func mustToken() string {
	token := flag.String("tg-token", "", "telegram bot token") // объявляем флан для получения токена при запуске программы
	
	flag.Parse()

	if *token == "" {
		log.Fatal("Отсутствует токен для бота")
	}

	return *token
}