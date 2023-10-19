package storage

type Storage interface {
	Add()
	Delete()
	All()
}

type Task struct {
	ID          uint64
	Title       string
	Description string
	Done        bool
}
