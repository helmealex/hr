package persistence

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	logger     *slog.Logger
	config     *DBConfig
	connection *gorm.DB
}

func New(
	logger *slog.Logger,
	config *DBConfig,
) (*DB, error) {
	pg, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: config.dsn(),
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	db := DB{
		logger:     logger,
		config:     config,
		connection: pg,
	}

	db.CreateTables()

	return &db, nil
}
