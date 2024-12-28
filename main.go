package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func Nats() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()

	// Publish a message to the "foo" subject
	err = nc.Publish("foo", []byte("Hello, NATS!"))
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return
	}
	fmt.Println("Message published to 'foo'")

	// Create a channel to receive messages
	msgChannel := make(chan *nats.Msg)

	// Subscribe to the "foo" subject and send messages to the channel
	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Println("Received message callback")
		msgChannel <- m // Send the message to the channel
	})
	if err != nil {
		fmt.Println("Error subscribing to 'foo':", err)
		return
	}

	// Wait for a message to be received
	select {
	case msg := <-msgChannel:
		fmt.Printf("Received message: %s\n", string(msg.Data))
	case <-time.After(5 * time.Second): // Timeout after 5 seconds if no message is received
		fmt.Println("Timeout: No message received")
	}
}

func main() {
	fmt.Println("start..")
	Nats()
	fmt.Println("end..")
}
