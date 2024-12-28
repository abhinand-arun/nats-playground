package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// MessageStruct defines the structure of the received message
type MessageStruct struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
		return
	}
	defer nc.Close()

	subject := "task.q"
	queueGroup := "workers" // Queue group name

	// Subscribe to the subject as part of a queue group
	_, err = nc.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) {
		var message MessageStruct

		// Deserialize the JSON message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			log.Printf("Failed to deserialize message: %v", err)
			return
		}

		// Process the message
		log.Printf("Received message: ID=%d, Content=%s, Time=%s", message.ID, message.Content, message.Time)
		fmt.Printf("Processed Message: %+v\n", message)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Printf("Subscribed to subject '%s' in queue group '%s'. Waiting for messages...", subject, queueGroup)

	// Keep the program running to listen for messages
	select {}
}
