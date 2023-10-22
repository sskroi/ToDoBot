package telegram

// Text for /help and /start cmd
const (
	helpMsg  = "/add - добавить задачу\n/close - завершить выполнение задачи\n/delete - удалить задачу\n/uncompl - посмотреть незавершенные задачи\n/compl - посмотреть завершенные задачи\n/help - список команд"
	startMsg = "/help для получения информации о боте"
)

const (
	incorrectTitleMsg   = "Некорректное название задачи.\n\nПопробуйте снова:"
	taskAlreadyExistMsg = "Задача с таким названием уже существует.\n\nПопробуйте другое название:"
	successTitleSetMsg  = "Название успешно установлено.\n\nВведите описание задачи:"
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
	addingMsg         = "Добавление задачи -> "
	addingTitleMsg    = "Введите уникальное название для новой задачи:"
	addingDescrMsg    = "Введите описание для новой задачи:"
	addingDeadlineMsg = "Введите дату дедлайна для новой задачи в формате ???:"
)
