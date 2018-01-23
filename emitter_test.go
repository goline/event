package event_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/goline/event"
)

var _ = Describe("FactoryEvent", func() {
	It("should allow to subscribe/unsubscribe/emit an event by name", func() {
		e := event.NewEvent("some-event", 5, false)
		emitter := event.NewEmitter()
		Expect(emitter.On("some-event", func(event event.Event) {
			Expect(event.Payload()).To(Equal(5))
		})).NotTo(BeNil())
		emitter.Emit(e)

		emitter.On("some-event", func(event event.Event) {
			Expect(event.Payload()).To(Equal(10))
		})
		emitter.Off("some-event")
		emitter.Emit(e) // should not raise error

		e = event.NewEvent("some-other-event", 5, false)
		emitter.Emit(e) // should not raise error
	})

	It("should allow to stop event's listeners in sequnce mode", func() {
		e := event.NewEvent("some-event", 5, false)
		emitter := event.NewEmitter()
		emitter.On("some-event", func(event event.Event) {
			event.WithPayload(10)
			event.Stop()
		})
		emitter.On("some-event", func(event event.Event) {
			event.WithPayload("some-string")
		})
		emitter.Emit(e)
		Expect(e.Payload()).To(Equal(10))
	})

	It("should allow to emit event's listeners in asynchronous mode", func() {
		e := event.NewEvent("some-event", 5, true)
		emitter := event.NewEmitter()
		emitter.On("some-event", func(event event.Event) {
			event.WithPayload(10)
		})
		emitter.On("some-event", func(event event.Event) {
			event.WithPayload("some-string")
		})
		emitter.Emit(e)
	})
})
