package producer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
)

type EventType = uint32

 const (
	Created EventType = iota
 	Updated
 	Removed
 )

 type Event struct {
 	Type EventType
 	Payload map[string]interface{}
 }

type Producer interface {
	Init(ctx context.Context) error
	SendEvent(event Event) error
	Close()
}

func NewProducer(broker string, topic string) Producer {
	return &producer{
		broker: broker,
		topic: topic,
	}
}

type producer struct {
	producer sarama.SyncProducer
	broker string
	topic string
}

func (pr *producer) Init(ctx context.Context) error {
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true

	client, err := sarama.NewClient([]string{pr.broker}, cfg)
	if err != nil {
		return err
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return err
	}

	pr.producer = producer
	return nil
}

func (pr *producer) SendEvent(event Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     pr.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(payload),
	}

	_, _, err = pr.producer.SendMessage(msg)
	return err
}

func (pr *producer) Close() {
	pr.producer.Close()
}
