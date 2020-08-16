package entity

type Job struct {
	JobName string            `json:"name"`
	JobArgs map[string]string `json:"args"`
}
