package domain

type EventRecorder struct {
	events DomainEvents
}

func InitEventRecorder() *EventRecorder {
	return &EventRecorder{events: make(DomainEvents, 0)}
}

func (ag *EventRecorder) Record(event DomainEvent) {
	ag.events = append(ag.events, event)
}

func (ag *EventRecorder) Flush() {
	ag.events = make(DomainEvents, 0)
}

func (ag *EventRecorder) Pull() DomainEvents {
	events := ag.events
	ag.Flush()

	return events
}
