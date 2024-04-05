package main

import (
	"log/slog"

	"github.com/xlzpm/internal/config"
	"github.com/xlzpm/pkg/logger/initlog"
)

func main() {
	cfg := config.MustConfig()

	log := initlog.InitLogger()

	log.Info("starting app", slog.Any("cfg", cfg))

	// TODO run server

	// TODO graceful shutdown
}
