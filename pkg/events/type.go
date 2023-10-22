package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type EvType int

// Types of events
const (
	Unknown EvType = iota
	Message
)

type Event struct {
	Type EvType
	Text string
	Meta interface{}
}