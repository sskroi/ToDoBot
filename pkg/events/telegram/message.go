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
	UnComplTasksMsg   = "Список незавершённых задач:\n\n"
	ComplTasks        = "Список завершённых задач:\n\n"
)

// Text for adding task
const (
	addingMsg            = "Добавление задачи -> "
	incorrectTitleMsg    = "Некорректное название задачи.\n\nПопробуйте снова:"
	taskAlreadyExistMsg  = "Задача с таким названием уже существует.\n\nПопробуйте другое название:"
	successTitleSetMsg   = "Название успешно установлено.\n\nВведите описание задачи:"
	addingTitleMsg       = "Введите уникальное название для новой задачи:"
	successDescrSetMsg   = "Описание задачи успешно установлено.\n\nВведите дату дедлайна для новой задачи в формате\n\n\"ДД-ММ-ГГГГ ЧЧ:ММ\":"
	incorrectDeadlineMsg = "Некорректный формат времени. Попробуйте снова.\n\nВведите дату дедлайна для новой задачи в формате\n\n\"ДД-ММ-ГГГГ ЧЧ:ММ\":"
	successDeadlineMsg   = "Задача успешно добавлена.\n\n/uncompl - для просмотра незавершённых задач."
)

func makeTasksString(tasks []storage.Task) string {
	dateTimeFormat := "02-01-2006 15:04"

	var res string = ""
	for _, v := range tasks {
		titleString := fmt.Sprintf("🧷 %s\n", v.Title)

		var statusString string
		if v.Done {
			statusString = fmt.Sprintf("✅ %s\n", getDoneStatus(v.Done))
		} else {
			statusString = fmt.Sprintf("❌ %s\n", getDoneStatus(v.Done))
		}

		deadlineString := fmt.Sprintf("❗️ Дедлайн: %s\n", time.Unix(int64(v.Deadline), 0).Format(dateTimeFormat))

		descrString := fmt.Sprintf("🗒 Описание: %s\n", v.Description)

		res += titleString + statusString + deadlineString + descrString + "\n"
	}

	return res
}

func getDoneStatus(status bool) string {
	if !status {
		return "Не выполнено"
	} else {
		return "Выполнено"
	}
}
