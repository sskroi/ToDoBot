package telegram

import (
	"ToDoBot1/pkg/clients/telegram"
)

type Processor struct {
	tg *telegram.Client
	offset int
}

func New() Processor