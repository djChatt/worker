package worker

import (
	"fmt"
	"log"
	"reflect"
	"worker/internal/pkg/entity"
)

type Worker struct {
	maxConcurrentRoutines int
	guard                 chan int
	jobChannel            chan entity.Job
	workList              map[string]*entity.Work
}

func NewWorker(maxRoutines int, jobChan chan entity.Job, worksList map[string]*entity.Work) *Worker {
	return &Worker{
		maxConcurrentRoutines: maxRoutines,
		guard:                 make(chan int, maxRoutines),
		jobChannel:            jobChan,
		workList:              worksList,
	}
}

func (w *Worker) Run() {
	for {
		w.guard <- 1
		job := <-w.jobChannel
		go w.Work(job)
	}
}

func (w *Worker) Work(job entity.Job) {
	fmt.Println("Recieved Job: ", job)
	work := w.workList[job.JobName]
	w.Execute(work, job.JobArgs)
	<-w.guard
}

func (w *Worker) Execute(work *entity.Work, jobArgs map[string]string) {
	workFunction := reflect.ValueOf(work.DynamicFunction)
	if len(jobArgs) != len(jobArgs) {
		log.Printf("expected number of args in %v is %v, recieved number of args is %v", work.WorkName, len(work.WorkArgs), len(jobArgs))
		return
	}
	in := make([]reflect.Value, workFunction.Type().NumIn())

	for i, argName := range work.WorkArgs {
		if _, exists := jobArgs[argName]; !exists {
			log.Printf("Missing arg with name: %v", argName)
			return
		}

		in[i] = reflect.ValueOf(jobArgs[argName])
	}
	log.Println(jobArgs)
	workFunction.Call(in)
}
