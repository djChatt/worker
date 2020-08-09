package worker

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

type Loader interface {
	Load(string, map[string]string) error
}

type loader struct {
	client    *redis.Client
	namespace string
}

func NewLoader(namespace string, client *redis.Client) Loader {
	return &loader{
		client:    client,
		namespace: namespace,
	}
}

func (l *loader) Load(functionName string, args map[string]string) error {
	jobEntry := map[string]interface{}{"func": functionName, "args": args}
	jsonEntry, err := json.Marshal(jobEntry)
	if err != nil {
		return err
	}

	err = l.client.RPush(l.namespace, string(jsonEntry)).Err()
	if err != nil {
		return err
	}
	return nil
}
