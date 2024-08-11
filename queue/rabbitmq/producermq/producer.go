package producermq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer interface {
	Publish(message string) error
	Close() error
}

type rabbitmqProducerImpl struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}


func NewRabbitMQProducer(url string) (RabbitMQProducer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close() 
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"learning_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &rabbitmqProducerImpl{
		conn:      conn,
		channel:   ch,
		queueName: q.Name,
	}, nil
}

// Publish xabarni RabbitMQ ga jo'natish
func (r *rabbitmqProducerImpl) Publish(message string) error {
	err := r.channel.Publish(
		"",
		r.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	return nil
}

// Close RabbitMQ bilan bog'lanishni yopish
func (r *rabbitmqProducerImpl) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	if err := r.conn.Close(); err != nil {
		return err
	}
	return nil
}
