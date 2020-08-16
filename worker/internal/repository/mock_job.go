package repository

import (
	"worker/internal/pkg/entity"

	"github.com/stretchr/testify/mock"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) GetJob(namespace string) (*entity.Job, error) {
	args := m.Called(namespace)
	return args.Get(0).(*entity.Job), args.Error(1)
}

func (m *MockJobRepository) CreateJob(namespace string, job entity.Job) error {
	args := m.Called(namespace, job)
	return args.Error(0)
}

func (m *MockJobRepository) ListAllJobs(namespace string) ([]entity.Job, error) {
	args := m.Called(namespace)
	return args.Get(0).([]entity.Job), args.Error(1)
}
