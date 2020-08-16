package loader

import (
	"testing"
	"worker/internal/pkg/entity"
	"worker/internal/repository"

	"github.com/stretchr/testify/assert"
)

//test if the data is being loaded in the database
func TestLoad(t *testing.T) {
	mockRepo := new(repository.MockJobRepository)
	actualJob := entity.Job{
		JobName: "test_job_name",
		JobArgs: map[string]string{"arg1": "val1", "arg2": "val2"},
	}
	mockRepo.On("CreateJob", "test_namespace", actualJob).Return(nil)
	testLoader := NewLoader("test_namespace", mockRepo)
	testError := testLoader.Load("test_job_name", map[string]string{"arg1": "val1", "arg2": "val2"})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, testError)
}

func TestListJobs(t *testing.T) {
	mockRepo := new(repository.MockJobRepository)
	actualJob := entity.Job{
		JobName: "test_job_name",
		JobArgs: map[string]string{"arg1": "val1", "arg2": "val2"},
	}
	mockRepo.On("ListAllJobs", "test_namespace").Return([]entity.Job{actualJob}, nil)
	testLoader := NewLoader("test_namespace", mockRepo)
	testRes, _ := testLoader.ListJobs()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, []entity.Job{actualJob}, testRes)
}
