package telegram

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/storage"
	"fmt"
	"time"
	"unicode/utf8"
)

// Text for cmds
const (
	helpCmd  = "/help"
	startCmd = "/start"
	notifCmd = "/notif"
)

// Text for /help and /start cmds
const (
	helpMsg = "🤖 Бот предоставляет реализацию простого 📌ToDo списка. Вы можете добавлять задачи, удалять, менять статус выполнения, а также просматривать список ваших задач.\n\n❗️ Вы можете * <code>cкопировать</code> * название задачи <code>кликнув</code> на него.\n\nВсё взаимодействие с ботом осуществляется с помощью интерактивного меню ⤵️"

	unknownCmdMsg = "❓ Неизвестная команда.\n\n/help - для просмотра доступных команд"
)

// Text for output information about tasks
const (
	noUncomplTasksMsg = "👌 У вас нет незавершённых задач."
	noComplTasksMsg   = "🤷🏻‍♀️ У вас нет завершённых задач."
	UnComplTasksMsg   = "⤵️ Список незавершённых задач:\n\n"
	ComplTasks        = "⤵️ Список завершённых задач:\n\n"
	taskNotExistMsg   = "❌ Задачи с таким названием не существует."
)

// Text for adding task
const (
	addingMsg            = "➕ Добавление задачи:\n\n"
	addingTitleMsg       = "📝 Введите уникальное название для новой задачи"
	incorrectTitleMsg    = "❌ Некорректное название задачи.\n🔄 Попробуйте снова"
	taskAlreadyExistMsg  = "❌ Задача с таким названием уже существует.\n🔄 Попробуйте другое название"
	successTitleSetMsg   = "✅ Название успешно установлено.\n\n📝 Введите описание задачи\n\n❗️ Если длина описания будет меньше двух символов, то описание не будет отображаться в списках задач."
	successDescrSetMsg   = "✅ Описание задачи успешно установлено.\n\n📝 Введите дату дедлайна для новой задачи в формате\n\n\"ДД.ММ.ГГГГ ЧЧ:ММ\"\n\n❗️ Вы можете не вводить время (ЧЧ:ММ), тогда оно будет установлено на <b>23:59</b>"
	incorrectDeadlineMsg = "❌ Некорректный формат времени.\n🔄 Попробуйте снова\n\n📝 Введите дату дедлайна для новой задачи в формате\n\n\"ДД.ММ.ГГГГ ЧЧ:ММ\"\n\n❗️ Вы можете не вводить время (ЧЧ:ММ), тогда оно будет установлено на <b>23:59</b>"
	successDeadlineMsg   = "✅ Задача успешно добавлена."
	TitleCantStartSlash  = "❌ Некорректное название задачи. Название не может начинаться с символа \"/\"\n🔄 Попробуйте снова"
	DescrCantStartSlash  = "❌ Некорректное описание задачи. Описание не может начинаться с символа \"/\"\n🔄 Попробуйте снова"
)

// Text for closing task
const (
	closingMsg              = "✔️ Завершение задачи:\n\n"
	closingTitleMsg         = "📝 Введите название задачи, которую вы хотите завершить"
	closingAlreadyClosedMsg = "☑️ Задача уже выполнена."
	closingSuccessClosed    = "✅ Задача успешно помечена как выполненная."
)

// Text for deleting task
const (
	deletingMsg           = "🗑 Удаление задачи:\n\n"
	deletingTitleMsg      = "📝 Введите название задачи, которую вы хотите удалить"
	deletingSuccessDelete = "✅ Задача успешно удалена."
)

// Text for main menu buttons
const (
	uncomplTasksBtn = "📌 My tasks"

	closeTaskBtn = "🏁 Complete"

	addTaskBtn = "➕ Add"
	delTaskBtn = "🗑 Delete"

	complTasksBtn = "🗄 Archive"
)

// reply markup keyboard main menu var
var mainMenuBtns = telegram.NewReplyKeyboard([][]string{
	{uncomplTasksBtn, closeTaskBtn},
	{addTaskBtn, delTaskBtn, complTasksBtn},
})

const (
	dateTimeFormat = "02.01.2006 15:04"
)

func UncomplTasksString(tasks []storage.Task) string {
	var res string = ""
	for _, v := range tasks {
		titleString := fmt.Sprintf("🔖 * <code>%s</code> *\n", v.Title)

		var timeToDeadLineStr string
		curTime := time.Now().Unix()
		if int64(v.Deadline) < curTime {
			diff := time.Unix(curTime, 0).Sub(time.Unix(int64(v.Deadline), 0))
			d := int(diff.Hours()) / 24
			h := int(diff.Hours()) % 24
			m := int(diff.Minutes()) % 60

			timeToDeadLineStr = fmt.Sprintf("🚫 Overdue: %dd %dh %dm\n", d, h, m)
		} else {
			diff := time.Unix(int64(v.Deadline), 0).Sub(time.Unix(curTime, 0))
			d := int(diff.Hours()) / 24
			h := int(diff.Hours()) % 24
			m := int(diff.Minutes()) % 60

			timeToDeadLineStr = fmt.Sprintf("⏳ Remaining: %dd %dh %dm\n", d, h, m)
		}

		deadlineString := fmt.Sprintf("🗓 Deadline: %s\n", time.Unix(int64(v.Deadline), 0).Format(dateTimeFormat))

		var descrString string
		if utf8.RuneCount([]byte(v.Description)) < 2 {
			descrString = ""
		} else {
			descrString = fmt.Sprintf("🗒 Description: %s\n", v.Description)
		}

		res += titleString + timeToDeadLineStr + deadlineString + descrString + "\n"
	}

	return res
}

func makeTasksString(tasks []storage.Task) string {
	var res string = ""
	for _, v := range tasks {
		titleString := fmt.Sprintf("🧷 * <code>%s</code> *\n", v.Title)

		var statusString string
		if v.Done {
			statusString = fmt.Sprintf("✅ %s\n", getDoneStatus(v.Done))
		} else {
			statusString = fmt.Sprintf("❌ %s\n", getDoneStatus(v.Done))
		}

		deadlineString := fmt.Sprintf("⏰ Дедлайн: %s\n", time.Unix(int64(v.Deadline), 0).Format(dateTimeFormat))

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
