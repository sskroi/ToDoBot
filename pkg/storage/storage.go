package storage

type Storage interface {
	Add(*Task) error
	Delete(*Task) error
	CloseTask(*Task) error
	UnCompl(*User) ([]Task, error)
	Compl(*User) ([]Task, error)
}

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
}
