package worker

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type JobRunner struct {
	namespace  string
	workList   map[string]*Work
	client     *redis.Client
	workers    []*Worker
	jobChannel chan string
	jobReader  Reader
}

func NewJobRunner(ns string, redisClient *redis.Client, numWorkers int, maxRoutinePerWorker int) *JobRunner {
	jobWorkers := make([]*Worker, numWorkers)
	jobChan := make(chan string)
	worksList := map[string]*Work{}
	for i := range jobWorkers {
		jobWorkers[i] = NewWorker(ns, maxRoutinePerWorker, redisClient, jobChan, worksList)
	}

	return &JobRunner{
		namespace:  ns,
		client:     redisClient,
		workList:   worksList,
		workers:    jobWorkers,
		jobChannel: jobChan,
		jobReader:  NewReader(redisClient, jobChan, ns),
	}
}

func (jr *JobRunner) RegisterWork(fn interface{}, workName string, args ...string) error {
	log.Println("starting register")
	log.Println(workName)
	if _, exists := jr.workList[workName]; exists {
		return errors.New("Already work exists by this name")
	}
	newWork, err := CreateWork(fn, workName, args...)
	if err != nil {
		return err
	}
	jr.workList[workName] = newWork
	return nil
}

func (jr *JobRunner) Start() {
	fmt.Println("Starting")
	for _, worker := range jr.workers {
		go worker.Run()
	}
	jr.jobReader.Read()
}
