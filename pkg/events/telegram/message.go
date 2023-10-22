package telegram

import (
	"ToDoBot1/pkg/storage"
	"fmt"
	"time"
	"unicode/utf8"
)

/*
list of commands for bot father^

add - Добавить задачу
close - Завершить выполнение задачи
delete - Удалить задачу
uncompl - Незавершённые задачи
compl - Завершённые задачи
help - Список всех команд
*/

// Text for /help and /start cmd
const (
	helpMsg     = "Бот предоставляет реализацию простого ToDo списка. Вы можете добавлять задачи, удалять, менять статус выполнения, а также просматривать список ваших задач.\n\nСписок команд бота:\n\n" + commandsMsg
	startMsg    = "/help для получения информации о боте"
	commandsMsg = "/add - добавить задачу\n/close - завершить выполнение задачи\n/delete - удалить задачу\n/uncompl - незавершённые задачи\n/compl - завёршенные задачи\n/help - список всех команд"

	unknownCmdMsg = "Неизвестная комманда. /help - для просмотра доступных команд."
)

// Text for output information about tasks
const (
	noUncomplTasksMsg = "У вас нет незавершённых задач."
	noComplTasksMsg   = "У вас нет завершённых задач."
	UnComplTasksMsg   = "Список незавершённых задач:\n\n"
	ComplTasks        = "Список завершённых задач:\n\n"
	taskNotExistMsg   = "❌ Задачи с таким названием не существует."
)

// Text for adding task
const (
	addingMsg            = "➕ Добавление задачи:\n\n"
	addingTitleMsg       = "Введите уникальное название для новой задачи."
	incorrectTitleMsg    = "Некорректное название задачи.\n\nПопробуйте снова."
	taskAlreadyExistMsg  = "Задача с таким названием уже существует.\n\nПопробуйте другое название."
	successTitleSetMsg   = "Название успешно установлено.\n\nВведите описание задачи.\n\nЕсли длина описания будет меньше двух символов, то описание не будет отображаться в списках задач."
	successDescrSetMsg   = "Описание задачи успешно установлено.\n\nВведите дату дедлайна для новой задачи в формате\n\n\"ДД-ММ-ГГГГ ЧЧ:ММ\""
	incorrectDeadlineMsg = "Некорректный формат времени.\n\nПопробуйте снова.\n\nВведите дату дедлайна для новой задачи в формате\n\n\"ДД-ММ-ГГГГ ЧЧ:ММ\""
	successDeadlineMsg   = "Задача успешно добавлена.\n\n/uncompl - для просмотра незавершённых задач."
)

// Text for closing task
const (
	closingMsg              = "✔️ Завершение задачи:\n\n"
	closingTitleMsg         = "Введите название задачи, которую вы хотите завершить."
	closingAlreadyClosedMsg = "Задача уже выполнена."
	closingSuccessClosed    = "✅ Задача успешно помечена как выполненная."
)

// Text for deleting task
const (
	deletingMsg           = "🗑 Удаление задачи:\n\n"
	deletingTitleMsg      = "Введите название задачи, которую вы хотите удалить."
	deletingSuccessDelete = "🗑 Задача успешно удалена."
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

		var descrString string
		if utf8.RuneCount([]byte(v.Description)) < 2 {
			descrString = ""
		} else {
			descrString = fmt.Sprintf("🗒 Описание: %s\n", v.Description)
		}

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
