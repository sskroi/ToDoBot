package main

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/config"
	"ToDoBot1/pkg/events/telegramproc"
	"ToDoBot1/pkg/handler"
	"ToDoBot1/pkg/storage/sqlite"
	"log"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	tgClient := telegram.New(config.Telegram)

    err := tgClient.SetWebhook(config.Server.URL, config.TLS.CertificatePath)
	if err != nil {
		log.Fatalf("can't set telegram webhook: %s", err.Error())
	}

	storage, err := sqlite.New(config.SQLite)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	processor := telegramproc.New(tgClient, storage)

    handler := handler.New(processor)

    http.HandleFunc("/", handler.HandleUpdate)
	err = http.ListenAndServeTLS(":"+config.Server.Port, config.TLS.CertificatePath, config.TLS.PrivateKeyPath, nil)
    if err != nil {
        log.Fatalf("can't start server: %s", err.Error())
    }
}
