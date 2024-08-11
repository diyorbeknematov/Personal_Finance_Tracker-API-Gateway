package producer

import (
	"context"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	SendMessage(ctx context.Context, topic string, key string, value string) error
	Close() error
}

type kafkaProducerImpl struct {
	writer *kafka.Writer
	logger *slog.Logger
}

func NewKafkaProducer(brokers []string, logger *slog.Logger) KafkaProducer {
	return &kafkaProducerImpl{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
			Balancer:               &kafka.LeastBytes{}, // Kamroq band bo'lgan partitsiyaga jo'natish
			BatchTimeout:           10 * time.Millisecond,
			ReadTimeout:            10 * time.Second, // O'qish vaqtini belgilash
			WriteTimeout:           10 * time.Second, // Yozish vaqtini belgilash
		},
		logger: logger,
	}
}

func (p *kafkaProducerImpl) SendMessage(ctx context.Context, topic string, key string, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Topic: topic,
	}

	retryCount := 3
	for i := 0; i < retryCount; i++ {
		err := p.writer.WriteMessages(ctx, msg)
		if err != nil {
			p.logger.Error("Error sending message to Kafka", "attempt", i+1, "error", err)
			if i < retryCount-1 {
				time.Sleep(2 * time.Second) // Qayta urinishdan oldin kutish
				continue
			}
			return err
		}
		p.logger.Info("Message sent successfully to Kafka", "topic", topic, "key", key)
		break
	}
	return nil
}

func (p *kafkaProducerImpl) Close() error {
	err := p.writer.Close()
	if err != nil {
		p.logger.Error("Error closing Kafka writer", "error", err)
	}
	return err
}
