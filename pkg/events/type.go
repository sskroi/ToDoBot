package events

type Fetcher interface {
	Fetcher(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type EvType int

const (
	Unknown EvType = iota
	Message
)

type Event struct {
	Type EvType
	Text string
}