package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"infrastructure/backend/service/gateway/internal/logging"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func Connect(config ConnectionConfig) (*amqp091.Connection, *amqp091.Channel, error) {
	connStr := "amqp://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/"

	conn, err := amqp091.Dial(connStr) // используем amqp091
	if err != nil {
		logging.Logger().Errorf("Failed to connect to RabbitMQ: %s", err)
		return nil, nil, err
	}

	channel, err := conn.Channel() // используем amqp091
	if err != nil {
		logging.Logger().Errorf("Failed to open a channel: %s", err)
		return nil, nil, err
	}

	return conn, channel, nil
}

func DeclareQueue(channel *amqp091.Channel, queueName string) (amqp091.Queue, error) {
	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logging.Logger().Errorf("Failed to declare a queue: %s", err)
		return amqp091.Queue{}, err
	}

	return queue, nil
}

func PublishMessage(channel *amqp091.Channel, queueName string, body []byte) error {
	err := channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		logging.Logger().Errorf("Failed to publish message to queue %s: %s", queueName, err)
		return err
	}

	logging.Logger().Infof("Message successfully published to queue: %s", queueName)
	return nil
}

func ConsumeMessages(channel *amqp091.Channel, queueName string) (<-chan amqp091.Delivery, error) {
	msgs, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logging.Logger().Errorf("Failed to consume messages from queue %s: %s", queueName, err)
		return nil, err
	}

	return msgs, nil
}