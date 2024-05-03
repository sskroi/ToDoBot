package telegramproc

import (
	"ToDoBot1/pkg/clients/telegram"
	"ToDoBot1/pkg/storage"
	"fmt"
	"log"
	"time"
)

// Text for cmds
const (
	helpCmd  = "/help"
	startCmd = "/start"
)

// Text for /help and /start cmds
const (
	helpMsg = "The bot provides an implementation of a simple ToDo list.\n" +
        "You can * <code>copy</code> * the task name by clicking on it.\n\n" +
		uncomplTasksBtn + "  →  list of unfinished tasks\n" +
		closeTaskBtn + "  →  mark the task as completed\n" +
		addTaskBtn + "  →  add new task\n" +
		delTaskBtn + "  →  delete task\n" +
		closeTaskBtn + "  →  list of completed tasks\n"

	unknownCmdMsg = "❓ Unknown command.\n\n/help - to view the available commands"
)

// Text for output information about tasks
const (
	noUncomplTasksMsg = "👌 You don't have any unfinished tasks.."
	noComplTasksMsg   = "🤷🏻‍♀️ You don't have completed tasks."
	UnComplTasksMsg   = "⤵️ List of uncompleted tasks:\n\n"
	ComplTasks        = "⤵️ List of completed tasks:\n\n"
	taskNotExistMsg   = "❌ There is no task with this name."
)

// Text for adding task
const (
	addingMsg            = "➕ Adding task:\n\n"
	addingTitleMsg       = "📝 Enter unique name for new task"
	incorrectTitleMsg    = "❌ Incorrect task name.\n🔄 Try again"
	taskAlreadyExistMsg  = "❌ Task with this name already exists.\n🔄 Try another name"
	successTitleSetMsg   = "✅ Name has been successfully set.\n\n📝 Enter deadline date for new task in the format\n\n\"dd.mm.YYYY HH:MM\"\n\nThe default value for time is <b>23:59</b>"
	incorrectDeadlineMsg = "❌ Incorrect time format.\n🔄 Try again\n\nEnter deadline date for new task in the format\n\n\"dd.mm.YYYY HH:MM\"\n\nThe default value for time is <b>23:59</b>"
	successDeadlineMsg   = "✅ Task was successfully added."
	TitleCantStartSlash  = "❌ Incorrect task name. The name cannot start with a character \"/\"\n🔄 Try again"
)

// Text for closing task
const (
	closingMsg              = "✔️ Completing task:\n\n"
	closingTitleMsg         = "📝 Enter the name of the task you want to complete"
	closingAlreadyClosedMsg = "☑️ Task has already been completed."
	closingSuccessClosed    = "✅ Task has been successfully marked as completed."
)

// Text for deleting task
const (
	deletingMsg           = "🗑 Deleting task:\n\n"
	deletingTitleMsg      = "📝 Enter the name of the task you want to delete"
	deletingSuccessDelete = "✅ Task was successfully deleted."
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

		res += titleString(v.Title) + timeToDeadLineStr + deadlineString(v.Deadline) + "\n"
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

		res += titleString(v.Title) + finishTimeStr + deadlineString(v.Deadline) + "\n"
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

