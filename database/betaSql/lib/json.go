package lib

import "encoding/json"

func ConvertToJson(item interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
