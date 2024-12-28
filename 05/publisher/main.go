package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Publish a message to the 'updates' subject
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Hello, message %d", i+1)

		// Publish message to "updates" subject
		err := nc.Publish("updates", []byte(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Published: %s\n", message)

		// Add a delay to simulate periodic publishing
		time.Sleep(1 * time.Second)
	}

	// Optionally flush the connection (ensure all messages are sent)
	nc.Flush()

	// Check if there was an error with flushing
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All messages published.")
}
