package reader

import (
	"fmt"
	"log"
	"time"
	"worker/internal/pkg/entity"
	"worker/internal/repository"
)

//TODO: add an error channel to stop the workers

type JobReader interface {
	ReadAndSendJobs()
	Read() (*entity.Job, error)
}

type Reader struct {
	jobRepository repository.JobRepository
	jobChannel    chan entity.Job
	namespace     string
}

func NewReader(jobRepo repository.JobRepository, jobChan chan entity.Job, ns string) JobReader {
	return &Reader{
		jobRepository: jobRepo,
		jobChannel:    jobChan,
		namespace:     ns,
	}
}

func (r *Reader) ReadAndSendJobs() {
	for {
		job, err := r.Read()
		if err != nil {
			if err.Error() == "redis: nil" {
				//log.Print("No jobs remaining, will check again after 1 second")
				time.Sleep(time.Second)
				continue
			}
			log.Println("Error reading from redis", err)
			continue
		}
		fmt.Println("Sending job")
		r.jobChannel <- *job
	}
}

func (r *Reader) Read() (*entity.Job, error) {
	job, err := r.jobRepository.GetJob(r.namespace)
	if err != nil {
		return nil, err
	}
	return job, nil
}
