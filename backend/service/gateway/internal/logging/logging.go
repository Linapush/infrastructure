package logging

import (
	"os"
	"golang.org/x/exp/slog"
)

func Init() {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "default-service"
	}

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	slog.Info("Logging initialized", slog.String("service", serviceName))
}