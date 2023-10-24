package processorloop

import (
	"ToDoBot1/pkg/events"
	"log"
	"time"
)

type ProcessorLoop struct {
	processor events.Processor
	batchSize int
}

func New(processor events.Processor, batchSize int) ProcessorLoop {
	return ProcessorLoop{
		processor: processor,
		batchSize: batchSize,
	}
}

func (p *ProcessorLoop) Start() error {
	for {
		gotEvents, err := p.processor.Fetch(p.batchSize)
		if err != nil {
			log.Printf("__ERR ProcessorLoop: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		err = p.handleEvents(gotEvents)
		if err != nil {
			log.Print(err)

			continue
		}

		time.Sleep(time.Millisecond * 50)
	}
}

func (p *ProcessorLoop) handleEvents(events []events.Event) error {
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
