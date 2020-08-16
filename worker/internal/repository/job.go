package repository

import (
	"encoding/json"
	"worker/internal/infra/database"
	"worker/internal/pkg/entity"
)

type JobRepository interface {
	GetJob(namespace string) (*entity.Job, error)
	CreateJob(namespace string, job entity.Job) error
	ListAllJobs(namespace string) ([]entity.Job, error)
}

type jobRepository struct {
	databaseClient database.DBClient
}

func NewJobRepository(dbClient database.DBClient) JobRepository {
	return &jobRepository{
		databaseClient: dbClient,
	}
}

func (jr *jobRepository) GetJob(namespace string) (*entity.Job, error) {
	jobString, err := jr.databaseClient.GetJob(namespace)
	if err != nil {
		return nil, err
	}
	var job entity.Job
	err = json.Unmarshal([]byte(jobString), &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (jr *jobRepository) CreateJob(namespace string, job entity.Job) error {
	jsonEntry, err := json.Marshal(&job)
	if err != nil {
		return err
	}

	err = jr.databaseClient.AddJob(namespace, string(jsonEntry))
	if err != nil {
		return err
	}
	return nil
}

func (jr *jobRepository) ListAllJobs(namespace string) ([]entity.Job, error) {
	jobListString, err := jr.databaseClient.ListAllJobs(namespace)
	if err != nil {
		return nil, err
	}
	jobList := []entity.Job{}
	var job entity.Job
	for _, jobString := range jobListString {
		err = json.Unmarshal([]byte(jobString), &job)
		if err != nil {
			return nil, err
		}
		jobList = append(jobList, job)
	}
	return jobList, nil
}
