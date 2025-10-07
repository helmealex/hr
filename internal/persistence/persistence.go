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

func (db *DB) GetJobs(withCandidates bool) ([]model.Job, error) {
	var jobs []model.Job
	tx := db.connection
	if withCandidates {
		tx = tx.Preload("Candidates") //TODO: replace with candidates count
	}
	if err := tx.Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (db *DB) GetJobByID(id uint) (*model.Job, error) {
	var job model.Job

	if err := db.connection.Preload("Candidates").First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (db *DB) CreateCandidate(
	candidate *model.Candidate,
) error {
	if err := db.connection.Create(candidate).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) GetCandidates() ([]model.Candidate, error) {
	var candidates []model.Candidate
	if err := db.connection.Find(&candidates).Error; err != nil {
		return nil, err
	}

	return candidates, nil
}

func (db *DB) GetCandidateByID(id uint) (*model.Candidate, error) {
	var candidate model.Candidate

	if err := db.connection.First(&candidate, id).Error; err != nil {
		return nil, err
	}

	return &candidate, nil
}

func (db *DB) GetVacancies() ([]model.VacancyReport, error) {
	vacancyResult := []model.VacancyReport{}

	err := db.connection.Model(&model.Job{}).
		Select("level, count(*) as count").
		Where("status = ?", "open").
		Group("level").
		Scan(&vacancyResult).Error
	if err != nil {
		return nil, err
	}

	return vacancyResult, nil
}

func (db *DB) GetApplicantsTotal() ([]model.ApplicantReport, error) {
	applicantResult := []model.ApplicantReport{}

	err := db.connection.Model(&model.Candidate{}).
		Select("job_id, jobs.title, count(*) as count").
		Joins("JOIN jobs ON candidates.job_id = jobs.id").
		Where("jobs.status = ?", "open").
		Group("job_id, jobs.title").
		Scan(&applicantResult).Error
	if err != nil {
		return nil, err
	}

	return applicantResult, nil
}
