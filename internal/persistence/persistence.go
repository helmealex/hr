package persistence

import (
	"hr/internal/persistence/model"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	logger     *slog.Logger
	config     *Config
	connection *gorm.DB
}

func New(
	logger *slog.Logger,
	config *Config,
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

func (db *DB) CreateJob(
	job *model.Job,
) error {
	if err := db.connection.Create(job).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateCandidate(
	candidate *model.Candidate,
) error {
	if err := db.connection.Create(candidate).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) GetJobs(withCandidates bool) ([]model.Job, error) {
	var jobs []model.Job
	tx := db.connection
	if withCandidates {
		tx = tx.Preload("Candidates")
	}
	if err := tx.Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (db *DB) GetCandidates() ([]model.Candidate, error) {
	var candidates []model.Candidate
	if err := db.connection.Find(&candidates).Error; err != nil {
		return nil, err
	}

	return candidates, nil
}
