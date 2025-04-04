package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "logs",
	})

	for i := 0; i < 10; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprintf("Hello Kafka %d", i)),
		}
		writer.WriteMessages(context.Background(), msg)
		fmt.Println("Produced:", string(msg.Value))
		time.Sleep(1 * time.Second)
	}

	writer.Close()
}
