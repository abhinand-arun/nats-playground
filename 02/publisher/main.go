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
	subject := "task.q"
	for i := 1; i < 10; i++ {
		count++
		message := fmt.Sprintf("Message %d: hello world", count)

		err = nc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published: %s", message)
		}

		fmt.Printf("message publish %d\n", count)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("all messages published!")
}
