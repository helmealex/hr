package main

import (
	"context"
	"hr/internal/config"
	"hr/internal/persistence"
	"log/slog"
	"os"
)

type App struct {
	db     *persistence.DB
	logger *slog.Logger
}

func main() {
	if err := Run(); err != nil {
		slog.Default().ErrorContext(context.TODO(), "application failed to start", "error", err)
	}
}

func Run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config, err := config.Load()
	if err != nil {
		logger.ErrorContext(context.TODO(), "failed to load configuration", "error", err)
		return err
	}

	_, err = persistence.New(
		logger,
		&config.DBConfig,
	)
	if err != nil {
		logger.ErrorContext(context.TODO(), "failed to initialize database", "error", err)
		return err
	}

	return nil
}
