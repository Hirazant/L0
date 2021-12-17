package nats

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func Connect() {
	sc, _ := stan.Connect("test-cluster", "ientIsdsD")

	// Simple Synchronous Publisher
	data := []byte("fdfdfd")
	sc.Publish("foo", data) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()
}
