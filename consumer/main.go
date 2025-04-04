package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "logs",
		GroupID: "log-group",
	})

	fmt.Println("Listening for messages...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Consumed: %s = %s\n", string(msg.Key), string(msg.Value))
	}
}
