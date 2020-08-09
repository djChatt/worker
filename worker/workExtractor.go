package worker

import "encoding/json"

func ExtractWork(workString string) (string, map[string]interface{}, error) {
	var obj map[string]interface{}
	json.Unmarshal([]byte(workString), &obj)
	workFunction := obj["func"].(string)
	workArgs := obj["args"].(map[string]interface{})
	return workFunction, workArgs, nil
}
