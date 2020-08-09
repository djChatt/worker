package worker

import (
	"errors"
	"log"
	"reflect"
)

type Work struct {
	WorkName        string
	WorkArgs        []string
	DynamicFunction reflect.Value
}

func CreateWork(fn interface{}, workName string, args ...string) (*Work, error) {
	workFunction := reflect.ValueOf(fn)
	if workFunction.Type().Kind().String() != "func" {
		return nil, errors.New("Input is not a function")
	}
	if workFunction.Type().NumIn() != len(args) {
		return nil, errors.New("Invalid number of arguments in contract")
	}
	work := Work{
		WorkName:        workName,
		WorkArgs:        args,
		DynamicFunction: workFunction,
	}
	return &work, nil
}

func (w Work) Execute(workArgs map[string]interface{}) {
	if len(workArgs) != len(w.WorkArgs) {
		log.Printf("expected number of args in %v is %v, recieved number of args is %v", w.WorkName, len(w.WorkArgs), len(workArgs))
		return
	}
	in := make([]reflect.Value, w.DynamicFunction.Type().NumIn())

	for i, argName := range w.WorkArgs {
		if _, exists := workArgs[argName]; !exists {
			log.Printf("Missing arg with name: %v", argName)
			return
		}

		in[i] = reflect.ValueOf(workArgs[argName])
	}
	log.Println(workArgs)
	w.DynamicFunction.Call(in)
}
