package mapping

import (
	"encoding/json"

	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

func Mapping[T any](data any) (*T, error) {
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return nil, errs.NewError(errs.ErrBadRequest, err)
	}

	var result T
	err = json.Unmarshal(dataJson, &result)
	if err != nil {
		return nil, errs.NewError(errs.ErrBadRequest, err)
	}

	return &result, nil
}
