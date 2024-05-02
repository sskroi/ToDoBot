package main

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/config"
	"ToDoBot1/pkg/events/telegramProc"
	"ToDoBot1/pkg/storage/sqlite"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	// creating an object for interacting with telegram api
	tgClient := telegram.New(config.Telegram)
	err := tgClient.DeleteWebhook()
	if err != nil {
		panic(err)
	}
	err = tgClient.SetWebhook(config.Server.URL, config.TLS.CertificatePath)
	if err != nil {
		panic(err)
	}

	// creating an object for interacting with sqlite3 storage
	storage, err := sqlite.New(config.SQLite)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	processor := telegramProc.New(tgClient, storage)

	http.ListenAndServeTLS(":"+config.Server.Port, config.TLS.CertificatePath, config.TLS.PrivateKeyPath,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            body, err := io.ReadAll(r.Body)
            if err != nil {
                log.Println("can't read telegram's answer: ", err.Error())
                return
            }

            update := telegram.Update{}       
            if err := json.Unmarshal(body, &update); err != nil {
                log.Println("can't parse telegram answer: ", err.Error())
                return
            }

            err = processor.Process(telegramProc.Event(update))
            if err != nil {
                log.Println("can't process event: ", err.Error())
                return
            }
		}))

	// mainLoop := processorloop.New(processor, 100)
	// err = mainLoop.Start()
	// if err != nil {
	// 	log.Fatalf("can't start main loop %s", err.Error())
	// }

}
