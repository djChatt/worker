package entity

import (
	"log"
	"reflect"
)

type Work struct {
	WorkName        string
	WorkArgs        []string
	DynamicFunction interface{}
}

func (w Work) Execute(workArgs map[string]string) {
	workFunction := reflect.ValueOf(w.DynamicFunction)
	if len(workArgs) != len(w.WorkArgs) {
		log.Printf("expected number of args in %v is %v, recieved number of args is %v", w.WorkName, len(w.WorkArgs), len(workArgs))
		return
	}
	in := make([]reflect.Value, workFunction.Type().NumIn())

	for i, argName := range w.WorkArgs {
		if _, exists := workArgs[argName]; !exists {
			log.Printf("Missing arg with name: %v", argName)
			return
		}

		in[i] = reflect.ValueOf(workArgs[argName])
	}
	log.Println(workArgs)
	workFunction.Call(in)
}
