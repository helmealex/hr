package persistence

import "hr/internal/persistence/model"

func (db *DB) CreateTables() error {
	return db.connection.AutoMigrate(
		&model.Candidate{},
	)
}
