package database

import (
	"github.com/go-redis/redis"
)

type DBClient interface {
	AddJob(namespace, value string) error
	GetJob(namespace string) (string, error)
	ListAllJobs(namespace string) ([]string, error)
}

type RedisDbClient struct {
	client *redis.Client
}

func NewRedisDbClient(redisClient *redis.Client) DBClient {
	return &RedisDbClient{
		client: redisClient,
	}
}

func (r *RedisDbClient) AddJob(namespace string, value string) error {
	err := r.client.RPush(namespace, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisDbClient) GetJob(namespace string) (string, error) {
	job, err := r.client.LPop(namespace).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			//log.Print("No jobs remaining, will check again after 1 second")
			return "", err
		}
		return "", err
	}
	return job, nil
}

func (r *RedisDbClient) ListAllJobs(namespace string) ([]string, error) {
	jobs, err := r.client.LRange(namespace, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
