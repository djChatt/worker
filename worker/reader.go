package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Reader struct {
	redisClient *redis.Client
	jobChannel  chan string
	namespace   string
}

func NewReader(rClient *redis.Client, jobChan chan string, ns string) Reader {
	return Reader{
		redisClient: rClient,
		jobChannel:  jobChan,
		namespace:   ns,
	}
}

func (r *Reader) Read() {
	for {
		job, err := r.redisClient.LPop(r.namespace).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				//log.Print("No jobs remaining, will check again after 1 second")
				time.Sleep(time.Second)
				continue
			}
			log.Println("Error reading from redis", err)
			continue
		}
		fmt.Println("Sending job")
		r.jobChannel <- job
	}
}
