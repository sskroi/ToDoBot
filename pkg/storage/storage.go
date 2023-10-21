package storage

type Storage interface {
	Add(*Task) error
	Delete(*Task) error
	CloseTask(*Task) error
	UnCompl(*User) ([]Task, error)
	Compl(*User) ([]Task, error)
}

// Types of state
const (
	DefState  int = 10
	Adding1   int = 21
	Adding2   int = 22
	Adding3   int = 23
	Deleting1 int = 31
	Closing1  int = 41
)

type Task struct {
	TaskId      uint64 `db:"task_id"`
	UserId      uint64 `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	CreateTime  uint64 `db:"create_time"`
	Deadline    uint64 `db:"deadline"`
	Done        bool   `db:"done"`
}

type User struct {
	UserId   uint64 `db:"user_id"`
	Username string `db:"username"`
	State    int    `db:"state"`
	CurTask  uint   `db:"cur_task"`
}
