package main

import (
	"fmt"
	"log"
	"time"
	"workerProject/worker"

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
	newLoader.Load("one_function", map[string]string{"number": "3"})
	newLoader.Load("one_function", map[string]string{"number": "1"})
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
