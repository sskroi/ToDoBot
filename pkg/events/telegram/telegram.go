package telegram

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/e"
	"ToDoBot1/pkg/events"
	"ToDoBot1/pkg/storage"
)

type Processor struct {
	tg      *telegram.Client
	storage storage.Storage
	offset  int
}

type Meta struct {
	UserId uint64
	ChatId uint64
	Date   uint64
}

func New(tgClient *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      tgClient,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get updates", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	result := make([]events.Event, 0, len(updates))

	for _, v := range updates {
		result = append(result, event(v))
	}

	p.offset = updates[len(updates)-1].UpdateId + 1

	return result, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			UserId: upd.Message.From.UserId,
			ChatId: upd.Message.Chat.ChatId,
			Date:   upd.Message.Date,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.EvType {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}
