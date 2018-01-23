package event_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/goline/event"
)

var _ = Describe("FactoryEvent", func() {
	It("should allow to get event's name", func() {
		e := event.NewEvent("some-event", nil, false)
		Expect(e.Name()).To(Equal("some-event"))
	})
	It("should allow to stop event", func() {
		e := event.NewEvent("some-event", nil, false)
		Expect(e.IsStopped()).To(BeFalse())
		e.Stop()
		Expect(e.IsStopped()).To(BeTrue())
	})
	It("should allow to set/get event's payload", func() {
		e := event.NewEvent("some-event", 5, false)
		Expect(e.Payload()).To(Equal(5))
		e.WithPayload("some-string")
		Expect(e.Payload()).To(Equal("some-string"))
	})
	It("should allow to get event's IsAsync", func() {
		e := event.NewEvent("some-event", nil, true)
		Expect(e.IsAsync()).To(BeTrue())
	})
})
