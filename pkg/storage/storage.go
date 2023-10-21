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
	Id          uint64
	UserId      uint64
	Title       string
	Description string
	CreateTime  uint64
	Deadline    uint64
	Done        bool
}

type User struct {
	Id       uint64
	Username string
	State    int
}
