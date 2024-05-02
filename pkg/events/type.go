package events

type Processor interface {
	Process(e Event) error
	Fetch(limit int) ([]Event, error)
    ProcessRequest(serializedUpdate []byte) error
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
