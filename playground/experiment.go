package main

import (
	"encoding/json"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://127.0.0.1:4222",
		nats.UserInfo("foo", "secret"),
	)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	nc.Subscribe("greeting", func(m *nats.Msg) {
		log.Printf("[Received] %s", string(m.Data))
	})

	payload := struct {
		RequestID string
		Data      []byte
		Timestamp int64
	}{
		RequestID: "experiment-1",
		Data:      []byte("Hello from Playground!"),
		Timestamp: 1234567890,
	}
	msg, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("[Publishing] %s", string(msg))
	nc.Publish("greeting", msg)
	runtime.Goexit()
}
