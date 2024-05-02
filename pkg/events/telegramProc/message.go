package telegramProc

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/storage"
	"fmt"
	"log"
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
	helpMsg = "🤖 Бот предоставляет реализацию простого 📌ToDo списка.\n" +
		"❗️ Вы можете * <code>cкопировать</code> * название задачи <code>кликнув</code> на него.\n\n" +
		uncomplTasksBtn + "  →  список незавершённых задач\n" +
		closeTaskBtn + "  →  отметить задачу как выполненную\n" +
		addTaskBtn + "  →  добавить новую задачу\n" +
		delTaskBtn + "  →  удалить задачу\n" +
		closeTaskBtn + "  →  список завершённых задач\n" +
		"\n" +
		notifCmd + "  →  команда для получения информации о ближайших дедлайнах. Вы можете отправить сообщение с этой командой по расписанию (schedule message) для получения уведомления в удобное вам время."

	unknownCmdMsg = "❓ Неизвестная команда.\n\n/help - для просмотра доступных команд"
)

// Text for output information about tasks
const (
	noUncomplTasksMsg = "👌 У вас нет незавершённых задач."
	noComplTasksMsg   = "🤷🏻‍♀️ У вас нет завершённых задач."
	UnComplTasksMsg   = "⤵️ List of uncompleted tasks:\n\n"
	ComplTasks        = "⤵️ List of completed tasks:\n\n"
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
	var res string
	for _, v := range tasks {
		var timeToDeadLineStr string
		curTime := time.Now().Unix()
		if int64(v.Deadline) < curTime {
			diff := time.Unix(curTime, 0).Sub(time.Unix(int64(v.Deadline), 0))
			d := int(diff.Hours()) / 24
			h := int(diff.Hours()) % 24
			m := int(diff.Minutes()) % 60

			timeToDeadLineStr = fmt.Sprintf("🚫 <b>%dd %dh %dm</b> overdue\n", d, h, m)

		} else {
			diff := time.Unix(int64(v.Deadline), 0).Sub(time.Unix(curTime, 0))
			d := int(diff.Hours()) / 24
			h := int(diff.Hours()) % 24
			m := int(diff.Minutes()) % 60

			timeToDeadLineStr = fmt.Sprintf("⏳ <b>%dd %dh %dm</b> remaining\n", d, h, m)
		}

		res += titleString(v.Title) + timeToDeadLineStr + deadlineString(v.Deadline) + descrString(v.Description) + "\n"
	}

	return res
}

func complTasksString(tasks []storage.Task) string {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("can't load time location") // bad
	}

	var res string
	for _, v := range tasks {
		finishTimeStr := fmt.Sprintf("⏱ <b>%s</b> finish time\n",
			time.Unix(int64(v.FinishTime), 0).In(location).Format(dateTimeFormat))

		res += titleString(v.Title) + finishTimeStr + deadlineString(v.Deadline) + descrString(v.Description) + "\n"
	}

	return res
}

func titleString(title string) string {
	titleString := fmt.Sprintf("🔖 * <code>%s</code> *\n", title)

	return titleString
}

func deadlineString(deadline uint64) string {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("can't load time location") // bad
	}

	deadlineString := fmt.Sprintf("🗓 <b>%s</b> deadline\n", time.Unix(int64(deadline),
		0).In(location).Format(dateTimeFormat))

	return deadlineString
}

func descrString(descr string) string {
	var descrString string
	if utf8.RuneCount([]byte(descr)) < 2 {
		descrString = ""
	} else {
		descrString = fmt.Sprintf("🧷 %s\n", descr)
	}

	return descrString
}
