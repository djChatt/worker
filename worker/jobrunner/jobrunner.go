package jobrunner

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"worker/internal/pkg/entity"
	"worker/internal/repository"
	"worker/jobRunner/reader"
	"worker/jobRunner/worker"
)

type JobRunner struct {
	namespace     string
	workList      map[string]*entity.Work
	jobRepository repository.JobRepository
	workers       []*worker.Worker
	jobChannel    chan entity.Job
	jobReader     reader.JobReader
}

func NewJobRunner(ns string, jobRepo repository.JobRepository, numWorkers int, maxRoutinePerWorker int) *JobRunner {
	jobWorkers := make([]*worker.Worker, numWorkers)
	jobChan := make(chan entity.Job)
	worksList := map[string]*entity.Work{}
	for i := range jobWorkers {
		jobWorkers[i] = worker.NewWorker(maxRoutinePerWorker, jobChan, worksList)
	}

	return &JobRunner{
		namespace:     ns,
		jobRepository: jobRepo,
		workList:      worksList,
		workers:       jobWorkers,
		jobChannel:    jobChan,
		jobReader:     reader.NewReader(jobRepo, jobChan, ns),
	}
}

func (jr *JobRunner) RegisterWork(fn interface{}, workName string, args ...string) error {
	log.Println("starting register")
	log.Println(workName)
	if _, exists := jr.workList[workName]; exists {
		return errors.New("Already work exists by this name")
	}
	err := isValidWork(fn, args...)
	if err != nil {
		return err
	}
	newWork := entity.Work{
		WorkName:        workName,
		WorkArgs:        args,
		DynamicFunction: fn,
	}
	jr.workList[workName] = &newWork
	return nil
}

func (jr *JobRunner) Start() {
	fmt.Println("Starting")
	for _, worker := range jr.workers {
		go worker.Run()
	}
	jr.jobReader.ReadAndSendJobs()
}

func (jr *JobRunner) GetRegisteredJobsList() map[string]*entity.Work {
	return jr.workList
}

func isValidWork(fn interface{}, args ...string) error {
	workFunction := reflect.ValueOf(fn)
	if workFunction.Type().Kind().String() != "func" {
		return errors.New("Input is not a function")
	}
	if workFunction.Type().NumIn() != len(args) {
		return errors.New("Invalid number of arguments in contract")
	}
	return nil
}
