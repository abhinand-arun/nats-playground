package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("cant connect to NATS: %v", err)
		return
	}

	defer nc.Close()

	var count int
	for {
		count++
		message := fmt.Sprintf("Message %d: hello world", count)
		nc.Publish("login", []byte(message))
		time.Sleep(1 * time.Second)
	}
}
