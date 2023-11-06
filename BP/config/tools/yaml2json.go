package tools

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

func Yaml2Json(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	temp := make(map[string]interface{})
	err = yaml.Unmarshal(data, &temp)
	if err != nil {
		return nil, err
	}
	return json.Marshal(temp)
}
