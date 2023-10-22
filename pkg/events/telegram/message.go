package telegram

import (
	"ToDoBot1/pkg/storage"
	"fmt"
	"time"
)

// Text for /help and /start cmd
const (
	helpMsg  = "/add - добавить задачу\n/close - завершить выполнение задачи\n/delete - удалить задачу\n/uncompl - посмотреть незавершенные задачи\n/compl - посмотреть завершенные задачи\n/help - список команд"
	startMsg = "/help для получения информации о боте"
)

// Text for output information about tasks
const (
	noUncomplTasksMsg = "У вас нет незавершённых задач."
	noComplTasksMsg   = "У вас нет завершённых задач."
	UnComplTasksMsg   = "Список незавершённых задач:\n"
	ComplTasks        = "Список завершённых задач:\n"
)

// Text for adding task
const (
	addingMsg           = "Добавление задачи -> "
	incorrectTitleMsg   = "Некорректное название задачи.\n\nПопробуйте снова:"
	taskAlreadyExistMsg = "Задача с таким названием уже существует.\n\nПопробуйте другое название:"
	successTitleSetMsg  = "Название успешно установлено.\n\nВведите описание задачи:"
	addingTitleMsg      = "Введите уникальное название для новой задачи:"
	successDescrSetMsg  = "Описание задачи успешно установлено.\n\nВведите дату дедлайна для новой задачи в формате ?:"
)

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
