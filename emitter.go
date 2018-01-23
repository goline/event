package event

import "sync"

type Emitter interface {
	// On subscribes an event listener
	On(name string, listener Listener) Emitter

	// Off unsubscribes an event by name
	Off(name string) Emitter

	// Emit fires an event
	Emit(event Event)
}

func NewEmitter() Emitter {
	return &FactoryEmitter{
		listeners: make(map[string][]Listener),
	}
}

type FactoryEmitter struct {
	listeners map[string][]Listener
}

func (e *FactoryEmitter) On(name string, listener Listener) Emitter {
	e.listeners[name] = append(e.listeners[name], listener)
	return e
}

func (e *FactoryEmitter) Off(name string) Emitter {
	e.listeners[name] = make([]Listener, 0)
	return e
}

func (e *FactoryEmitter) Emit(event Event) {
	listeners, ok := e.listeners[event.Name()]
	if !ok {
		return
	}

	if event.IsAsync() {
		e.runAsync(event, listeners)
	} else {
		e.runSequence(event, listeners)
	}
}

func (e *FactoryEmitter) runSequence(event Event, listeners []Listener) {
	for _, listener := range listeners {
		if event.IsStopped() {
			break
		}
		listener(event)
	}
}

func (e *FactoryEmitter) runAsync(event Event, listeners []Listener) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for _, listener := range listeners {
		wg.Add(1)
		go func(listener Listener) {
			defer wg.Done()
			mutex.Lock()
			listener(event)
			mutex.Unlock()
		}(listener)
	}
	wg.Wait()
}
