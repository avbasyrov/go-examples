package main

import (
	"fmt"
	"go-examples/kafka/internal/kafka"
	"log"
	"strconv"
	"time"
)

func main() {
	const (
		kafkaHost = "kafka"
		topicName = "myTopic"
	)

	fmt.Println("starting...")

	time.Sleep(5 * time.Second)
	fmt.Println("\n\n run consumer group")
	consumerGroup, err := kafka.NewConsumer(kafkaHost, "myGroup-1")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(consumerGroup *kafka.Group) {
		err := consumerGroup.Close()
		if err != nil {
			log.Println(err)
		}
	}(consumerGroup)

	// Подписываемся на чтение топиков
	err = consumerGroup.Subscribe([]string{topicName})
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("\n\n run producer")
	producer, err := kafka.NewProducer(kafkaHost)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Close()

	go func() {
		for i := 1; i <= 10; i++ {
			message := []byte("Сообщение в кафку №" + strconv.Itoa(i))
			err := producer.SendMessage(topicName, message)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
	// Wait for message deliveries before shutting down
	defer producer.Flush(15 * 1000)

	// Читаем сообщения из Кафки
	for {
		msg, err := consumerGroup.ReadMessage()
		if err == nil {
			fmt.Printf("Получили сообщение из %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
