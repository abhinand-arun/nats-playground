package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
		return
	}
	defer nc.Close()

	subject := "task.q"
	queueGroup := "workers" // Queue group name

	fmt.Println("lis >>>.")

	// Subscribe to the subject as part of a queue group
	_, err = nc.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Printf("Subscribed to subject '%s' in queue group '%s'. Waiting for messages...", subject, queueGroup)

	// Keep the program running to listen for messages
	select {}
}
