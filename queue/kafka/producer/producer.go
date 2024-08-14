package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	ProducerMessage(topic string, msg []byte) error
	Close()
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string) KafkaProducer {
	log.Println("New kafka producer initialized with brokers:", brokers)
	return &kafkaProducer{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			AllowAutoTopicCreation: true,
		},
	}
}

func (p *kafkaProducer) ProducerMessage(topic string, msg []byte) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: msg,
	})
}

func (p *kafkaProducer) Close() {
	p.writer.Close()
}
