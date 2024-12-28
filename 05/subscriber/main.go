package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create a NATS subscription with acknowledgement
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}

	// Loop to listen for messages and acknowledge them after processing
	for {
		msg, err := sub.NextMsg(0)
		if err != nil {
			log.Fatal(err)
		}

		// Process the message (simulated with a print statement)
		fmt.Printf("Received message: %s\n", string(msg.Data))

		// Send acknowledgment that the message has been processed
		msg.Ack()

		// Optionally, you can log or handle failed ack scenarios if needed
	}
}
