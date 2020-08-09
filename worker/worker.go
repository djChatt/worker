package worker

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type Worker struct {
	namespace             string
	maxConcurrentRoutines int
	guard                 chan int
	redisClient           *redis.Client
	jobChannel            chan string
	workList              map[string]*Work
}

func NewWorker(ns string, maxRoutines int, rClient *redis.Client, jobChan chan string, worksList map[string]*Work) *Worker {
	return &Worker{
		namespace:             ns,
		maxConcurrentRoutines: maxRoutines,
		guard:                 make(chan int, maxRoutines),
		redisClient:           rClient,
		jobChannel:            jobChan,
		workList:              worksList,
	}
}

func (w *Worker) Run() {
	for {
		w.guard <- 1
		jobDetails := <-w.jobChannel
		go w.Work(jobDetails)
	}
}

func (w *Worker) Work(jobDetails string) {
	fmt.Println("Recieved Job: ", jobDetails)
	workFuncName, workArgs, err := ExtractWork(jobDetails)
	if err != nil {
		log.Print(err)
	}
	work := w.workList[workFuncName]
	work.Execute(workArgs)
	<-w.guard
}
