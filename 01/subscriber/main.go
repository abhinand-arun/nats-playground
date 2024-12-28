package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// Asynchronous Subscription
func NatsSubscribe(nc *nats.Conn) {
	fmt.Println("Starting asynchronous subscription...")
	_, err := nc.Subscribe("login", func(m *nats.Msg) {
		fmt.Printf("I got the message asynchronously: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println("Error in asynchronous subscription:", err)
		return
	}
}

// Synchronous Subscription
func NatsSubscribeSync(nc *nats.Conn) {
	fmt.Println("Starting synchronous subscription...")
	sub, err := nc.SubscribeSync("login")
	if err != nil {
		fmt.Println("Error in synchronous subscription:", err)
		return
	}

	// Consume messages synchronously
	go func() {
		for {
			msg, err := sub.NextMsg(20 * time.Second) // Wait for up to 20 sec for a message
			if err != nil {
				fmt.Println("Error receiving message:", err)
				break
			}
			fmt.Printf("I got the message synchronously: %s\n", string(msg.Data))
		}
	}()
}

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()

	// Testing asynchronous subscription
	//NatsSubscribe(nc)

	// Testing synchronous subscription
	NatsSubscribeSync(nc)

	// Keep the program running to process messages
	time.Sleep(1 * time.Minute)
}
