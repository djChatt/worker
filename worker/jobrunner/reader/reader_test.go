package reader

import (
	"testing"
	"worker/internal/pkg/entity"
	"worker/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	mockRepo := new(repository.MockJobRepository)
	actualJob := &entity.Job{
		JobName: "test_job_name",
		JobArgs: map[string]string{"arg1": "val1", "arg2": "val2"},
	}
	mockRepo.On("GetJob", "test_namespace").Return(actualJob, nil)
	testJobChan := make(chan entity.Job)
	testReader := NewReader(mockRepo, testJobChan, "test_namespace")
	testResultJob, testError := testReader.Read()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, actualJob, testResultJob)
	assert.Nil(t, testError)
}
