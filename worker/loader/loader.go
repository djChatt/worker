package loader

import (
	"worker/internal/pkg/entity"
	"worker/internal/repository"
)

type Loader interface {
	Load(string, map[string]string) error
	ListJobs() ([]entity.Job, error)
}

type loader struct {
	jobRepository repository.JobRepository
	namespace     string
}

func NewLoader(namespace string, jobRepo repository.JobRepository) Loader {
	return &loader{
		jobRepository: jobRepo,
		namespace:     namespace,
	}
}

func (l *loader) Load(functionName string, args map[string]string) error {
	job := entity.Job{
		JobName: functionName,
		JobArgs: args,
	}
	return l.jobRepository.CreateJob(l.namespace, job)
}

func (l *loader) ListJobs() ([]entity.Job, error) {
	return l.jobRepository.ListAllJobs(l.namespace)
}
