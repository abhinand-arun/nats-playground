package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// MessageStruct defines the structure of the message to be published
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
	for i := 1; i <= 10; i++ {
		// Create a message instance
		message := MessageStruct{
			ID:      i,
			Content: fmt.Sprintf("Message %d: hello world", i),
			Time:    time.Now().Format(time.RFC3339),
		}

		// Serialize the message to JSON
		messageData, err := json.Marshal(message)
		if err != nil {
			log.Printf("Failed to serialize message: %v", err)
			continue
		}

		// Publish the message
		err = nc.Publish(subject, messageData)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published: %s", message.Content)
		}

		fmt.Printf("Message published: %d\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All messages published!")
}
