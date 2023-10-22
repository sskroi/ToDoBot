package telegram

import (
	"ToDoBot1/pkg/e"
	"ToDoBot1/pkg/storage"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	HelpCmd    = "/help"
	StartCmd   = "/start"
	AddCmd     = "/add"
	CloseCmd   = "/close"
	UncomplCmd = "/uncompl"
	ComplCmd   = "/compl"
	DelCmd     = "/delete"
)

func (p *Processor) handleMsg(text string, meta Meta) error {
	text = strings.TrimSpace(text)

	log.Printf("got new messgae: %s | from: %s", text, meta.Username)

	userState, err := p.storage.GetState(meta.UserId)
	if err != nil {
		return e.Wrap("can't handle message", err)
	}

	switch userState {
	case storage.DefState:
		err = p.doCmd(text, meta)
	case storage.Adding1:
		err = p.adding1(text, meta)
	case storage.Adding2:
		err = p.adding2(text, meta)
	}

	if err != nil {
		return e.Wrap("can't handle message", err)
	}

	return nil
}

func (p *Processor) doCmd(text string, meta Meta) error {
	var err error

	switch text {
	case StartCmd:
		err = p.doStartCmd(meta)
	case HelpCmd:
		err = p.doHelpCmd(meta)
	case AddCmd:
		err = p.doAddCmd(meta)
	case CloseCmd:
		err = p.doCloseCmd(meta)
	case UncomplCmd:
		err = p.doUncomplCmd(meta)
	case ComplCmd:
		err = p.doComplCmd(meta)
	default:
		err = p.doUnknownCmd(meta)
	}
	if err != nil {
		return e.Wrap("can't do cmd", err)
	}

	return nil
}

func (p *Processor) doUnknownCmd(meta Meta) error {
	err := p.tg.SendMessage(int(meta.ChatId), "Неизвестная комманда. /help - для просмотра доступных команд.")
	if err != nil {
		return e.Wrap("can't do UnknownCmd", err)
	}

	return nil
}

func (p *Processor) doStartCmd(meta Meta) error {
	err := p.tg.SendMessage(int(meta.ChatId), startMsg)
	if err != nil {
		return e.Wrap("can't do /start", err)
	}

	return nil
}

func (p *Processor) doHelpCmd(meta Meta) error {
	err := p.tg.SendMessage(int(meta.ChatId), helpMsg)
	if err != nil {
		return e.Wrap("can't do /help", err)
	}

	return nil
}

func (p *Processor) doAddCmd(meta Meta) error {
	err := p.storage.Add(meta.UserId)
	if err != nil {
		return e.Wrap("can't do /add", err)
	}

	err = p.storage.SetState(meta.UserId, storage.Adding1)
	if err != nil {
		return e.Wrap("can't do /add", err)
	}

	err = p.tg.SendMessage(int(meta.ChatId), "Добавление задачи -> введите уникальное название:")
	if err != nil {
		return e.Wrap("can't do /add", err)
	}

	return nil
}

func (p *Processor) doCloseCmd(meta Meta) error {
	err := p.storage.SetState(meta.UserId, storage.Closing1)
	if err != nil {
		return e.Wrap("can't do /close", err)
	}

	err = p.tg.SendMessage(int(meta.ChatId), "Завершение задачи -> введите название задачи:")
	if err != nil {
		return e.Wrap("can't do /add", err)
	}

	return nil
}

func (p *Processor) doUncomplCmd(meta Meta) error {
	tasks, err := p.storage.Uncompl(meta.UserId)
	if err != nil {
		return e.Wrap("can't do /uncompl", err)
	}

	if len(tasks) == 0 {
		p.tg.SendMessage(int(meta.ChatId), noUncomplTasksMsg)
		if err != nil {
			return e.Wrap("can't do /uncompl", err)
		}

		return nil
	}

	tasksStr := makeTasksString(tasks)

	sentStr := UnComplTasksMsg + tasksStr

	p.tg.SendMessage(int(meta.ChatId), sentStr)
	if err != nil {
		return e.Wrap("can't do /uncompl", err)
	}

	return nil
}

func (p *Processor) doComplCmd(meta Meta) error {
	tasks, err := p.storage.Compl(meta.UserId)
	if err != nil {
		return e.Wrap("can't do /compl", err)
	}

	if len(tasks) == 0 {
		p.tg.SendMessage(int(meta.ChatId), noComplTasksMsg)
		if err != nil {
			return e.Wrap("can't do /uncompl", err)
		}

		return nil
	}

	tasksStr := makeTasksString(tasks)

	sentStr := ComplTasks + tasksStr

	p.tg.SendMessage(int(meta.ChatId), sentStr)
	if err != nil {
		return e.Wrap("can't do /compl", err)
	}

	return nil
}

func (p *Processor) adding1(text string, meta Meta) error {
	const errMsg = "can't add title to task"

	if text == "" {
		err := p.tg.SendMessage(int(meta.ChatId), addingMsg+incorrectTitleMsg)
		if err != nil {
			return e.Wrap(errMsg, err)
		}

		return nil
	}

	err := p.storage.UpdTitle(meta.UserId, text)
	if errors.Is(err, storage.ErrUnique) {
		if err := p.tg.SendMessage(int(meta.ChatId), addingMsg+taskAlreadyExistMsg); err != nil {
			return e.Wrap(errMsg, err)
		}

		return nil
	} else if err != nil {
		return e.Wrap(errMsg, err)
	}

	if err := p.storage.SetState(meta.UserId, storage.Adding2); err != nil {
		return e.Wrap(errMsg, err)
	}

	if err := p.tg.SendMessage(int(meta.ChatId), addingMsg+successTitleSetMsg); err != nil {
		return e.Wrap(errMsg, err)
	}

	return nil
}

func (p *Processor) adding2(text string, meta Meta) error {
	err := p.storage.UpdDescription(meta.UserId, text)
	if err != nil {
		return e.Wrap("can't add description to task", err)
	}

	if err := p.storage.SetState(meta.UserId, storage.Adding3); err != nil {
		return e.Wrap("can't add description to task", err)
	}

	if err := p.tg.SendMessage(int(meta.ChatId), successDescrSetMsg); err != nil {
		return e.Wrap("can't add description to task", err)
	}

	return nil
}

func makeTasksString(tasks []storage.Task) string {
	var res string = ""
	for _, v := range tasks {
		res += fmt.Sprintf("%s | Статус: %s | Описание: %s | Дедлайн: %s\n", v.Title, getDoneStatus(v.Done), v.Description, time.Unix(int64(v.Deadline), 0))
	}

	return res
}

func getDoneStatus(status bool) string {
	if !status {
		return "не завершена"
	} else {
		return "завершена"
	}
}
