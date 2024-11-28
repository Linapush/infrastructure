package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	slog.Info("Health check passed")

	return c.SendString("Service is healthy")
}
