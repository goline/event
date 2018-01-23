package event

type Event interface {
	// Name returns name of event
	Name() string

	// IsStopped returns true if event is stopped
	IsStopped() bool

	// IsAsync returns true if event is an asynchronous event
	IsAsync() bool

	// Stop stops event
	// Stop will not work under Async
	Stop()

	// Payload returns event's payload
	Payload() interface{}

	// WithPayload allows to set event's payload
	WithPayload(payload interface{}) Event
}

func NewEvent(name string, payload interface{}, isAsync bool) Event {
	return &FactoryEvent{
		name:      name,
		isStopped: false,
		isAsync:   isAsync,
		payload:   payload,
	}
}

type FactoryEvent struct {
	name      string
	isStopped bool
	isAsync   bool
	payload   interface{}
}

func (e *FactoryEvent) Name() string {
	return e.name
}

func (e *FactoryEvent) IsStopped() bool {
	return e.isStopped
}

func (e *FactoryEvent) IsAsync() bool {
	return e.isAsync
}

func (e *FactoryEvent) Stop() {
	if !e.isStopped {
		e.isStopped = true
	}
}

func (e *FactoryEvent) Payload() interface{} {
	return e.payload
}

func (e *FactoryEvent) WithPayload(payload interface{}) Event {
	e.payload = payload
	return e
}
