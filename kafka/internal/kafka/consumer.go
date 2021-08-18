package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Group struct {
	consumer  *kafka.Consumer
	kafkaHost string
	groupName string
}

func NewConsumer(kafkaHost, groupName string) (*Group, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
		"group.id":          groupName,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &Group{consumer: c}, nil
}

func (g *Group) Subscribe(topics []string) error {
	return g.consumer.SubscribeTopics(topics, nil)
}

func (g *Group) ReadMessage() (*kafka.Message, error) {
	return g.consumer.ReadMessage(-1)
}

func (g *Group) Close() error {
	return g.consumer.Close()
}
