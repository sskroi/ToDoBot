package processorloop

import (
	"ToDoBot1/pkg/events"
	"log"
	"time"
)

type ProcLoop struct {
	processor events.Processor
	batchSize int
}

func New(processor events.Processor, batchSize int) ProcLoop {
	return ProcLoop{
		processor: processor,
		batchSize: batchSize,
	}
}

func (p *ProcLoop) Start() error {
	for {
		gotEvents, err := p.processor.Fetch(p.batchSize)
		if err != nil {
			log.Printf("ERROR ProcLoop: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(time.Millisecond * 10)

			continue
		}

		err = p.handleEvents(gotEvents)
		if err != nil {
			log.Print(err)

			continue
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func (p *ProcLoop) handleEvents(events []events.Event) error {
	for _, event := range events {
		// log.Printf("got new event: %s", event.Text)

		err := p.processor.Process(event)
		if err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
