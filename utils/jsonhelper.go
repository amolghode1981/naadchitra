package utils

import "encoding/json"

func ToJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}

func ToPrettyJSON(v any) (string, error) {
	data, error := json.MarshalIndent(v, "", "  ")
	if error == nil {
		return string(data), nil
	}
	return "", error
}

func FromJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
