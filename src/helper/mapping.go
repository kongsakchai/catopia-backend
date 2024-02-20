package helper

import "encoding/json"

func Mapping[T any](data any) (*T, error) {
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(dataJson, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
