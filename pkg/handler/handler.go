package handler

import (
	"ToDoBot1/pkg/events"
	"io"
	"log"
	"net/http"
)

type Handler struct {
    processor events.Processor
}

func New(processor events.Processor) *Handler {
    return &Handler{
        processor: processor,
    }
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Println("can't read request body: ", err.Error())
        return
    }
    r.Body.Close()

    err = h.processor.ProcessRequest(body)
    if err != nil {
        log.Println("can't process request: ", err.Error())
        return
    }
}
