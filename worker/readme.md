# Worker
Worker is a Go library dealing with executing jobs in the background. You can use it to load jobs and run them in the background. The jobs are stored in a Redis cluster and can be persisted.

# How to use

## To load new Jobs in the queue
Make a loader with a namespace and a redis client to load jobs. The arguments are given as a map with string key/value pair. And the name of the function is given as a string. Before running the functions has to be registered in the same namespace. 
```go
package main

import (

	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	newLoader := worker.NewLoader("default", redisClient)
	err := newLoader.Load("one_function", map[string]string{"number": "2"})
	if err != nil {
		log.Fatal(err)
    }
}
```

## To Run jobs in the queue
Make a JobRunner with a namespace you want and a redis client. Register the functions with the function name as a string and a contract of the argument names. After that start the JobRunner.
```go
package main

import (

	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	jr := worker.NewJobRunner("default", redisClient, 1, 1)
	err = jr.RegisterWork(oneFunction, "one_function", "number")
	if err != nil {
		log.Fatal(err)
	}
	jr.Start()

}

func oneFunction(number string) {
	fmt.Println("function called with value", number)
	time.Sleep(2 * time.Second)
}
``` 