package handlers

import (
	"github.com/gofiber/fiber/v2"
	"infrastructure/backend/service/gateway/internal/logging"
	"infrastructure/backend/service/gateway/internal/rabbitmq"
)

func RpcHandler(c *fiber.Ctx) error {
	config := rabbitmq.ConnectionConfig{
		Host:     "rabbitmq",
		Port:     "5672",
		Username: "guest",
		Password: "guest",
	}

	conn, channel, err := rabbitmq.Connect(config)
	if err != nil {
		return c.Status(500).SendString("Failed to connect to RabbitMQ")
	}
	defer conn.Close()
	defer channel.Close()

	message := c.Body()

	err = rabbitmq.PublishMessage(channel, "rpc_queue", message)
	if err != nil {
		return c.Status(500).SendString("Failed to process RPC")
	}

	logging.Logger().Infof("Message published: %s", string(message))

	return c.SendString("RPC request processed")
}