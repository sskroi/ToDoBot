package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int              `json:"update_id"`
	Message  *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Date uint64 `json:"date"`
}

type From struct {
	UserId   uint64 `json:"id"`
	Username string `json:"username"`
}

type Chat struct {
	ChatId uint64 `json:"id"`
}

var ReplyKeyboardRemove = &struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
}{
	true,
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
}

func NewReplyKeyboard(buttons [][]string) *ReplyKeyboardMarkup {
	return &ReplyKeyboardMarkup{
		Keyboard:        buttons,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
}
