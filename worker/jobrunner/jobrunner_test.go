package jobrunner

import (
	"fmt"
	"reflect"
	"testing"
	"worker/internal/pkg/entity"
	"worker/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestRegisterWork(t *testing.T) {
	mockRepo := new(repository.MockJobRepository)
	testJobRunner := NewJobRunner("test-default", mockRepo, 1, 1)

	testFunction := func(arg1, arg2 string) string {
		return fmt.Sprint(arg1, arg2)
	}

	err := testJobRunner.RegisterWork(testFunction, "test_name", "arg1", "arg2")
	assert.Nil(t, err)

	actualWork := entity.Work{
		WorkName:        "test_name",
		WorkArgs:        []string{"arg1", "arg2"},
		DynamicFunction: testFunction,
	}
	testWork, ok := testJobRunner.GetRegisteredJobsList()["test_name"]

	assert.Equal(t, true, ok)
	assert.Equal(t, actualWork.WorkName, testWork.WorkName)
	assert.Equal(t, actualWork.WorkArgs, testWork.WorkArgs)
	assert.Equal(t, reflect.ValueOf(actualWork.DynamicFunction), reflect.ValueOf(testWork.DynamicFunction))

	err = testJobRunner.RegisterWork(testFunction, "test_name", "arg1", "arg2")
	assert.EqualError(t, err, "Already work exists by this name")
}

func TestRegisterWorkInvalidArgumentNumber(t *testing.T) {
	mockRepo := new(repository.MockJobRepository)
	testJobRunner := NewJobRunner("test-default", mockRepo, 1, 1)

	testFunction := func(arg1, arg2 string) string {
		return fmt.Sprint(arg1, arg2)
	}

	err := testJobRunner.RegisterWork(testFunction, "test_name", "arg1")
	assert.EqualError(t, err, "Invalid number of arguments in contract")
	_, ok := testJobRunner.GetRegisteredJobsList()["test_name"]
	assert.Equal(t, false, ok)
}
