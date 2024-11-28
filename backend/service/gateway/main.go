package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"infrastructure/backend/service/gateway/handlers"
	"infrastructure/backend/service/gateway/internal/logging"
)

func main() {
	logging.Init()

	app := fiber.New()

	app.Get("/health", handlers.HealthCheckHandler)
	app.Post("/rpc", handlers.RpcHandler)           

	log.Println("Service started on :8081")

	if err := app.Listen(":8081"); err != nil {
		log.Fatalf("Failed to start service: %s", err)
	}
}