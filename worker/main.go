package main

import (
	"fmt"
	"time"
	"worker/internal/pkg/entity"
)

func main() {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
	// databaseClient := database.NewRedisDbClient(redisClient)
	// jobRepo := repository.NewJobRepository(databaseClient)
	// newLoader := loader.NewLoader("default", jobRepo)
	// err := newLoader.Load("one_function", map[string]string{"number": "2"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newLoader.Load("one_function", map[string]string{"number": "3"})
	// newLoader.Load("one_function", map[string]string{"number": "1"})
	// jr := jobrunner.NewJobRunner("default", jobRepo, 1, 1)
	// err = jr.RegisterWork(oneFunction, "one_function", "number")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jr.Start()
	newWork, _ := entity.CreateWork(oneFunction, "one_function", "number")
}

func oneFunction(number string) {
	fmt.Println("function called with value", number)
	time.Sleep(2 * time.Second)
}
