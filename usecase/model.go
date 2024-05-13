package usecase

import (
	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/utils/onnx"
)

var min []float64 = []float64{-0.96437006, -2.43380029, -4.06119942}
var max []float64 = []float64{2.6400001, 2.43380022, 4.06119942}

type ModelUsecae struct {
	catModel  onnx.Model
	pcaModel  onnx.Model
	stdModel  onnx.Model
	userModel onnx.Model
}

func NewModelUsecae() domain.ModelUsecae {
	return &ModelUsecae{
		catModel:  onnx.CreateModel("cat"),
		pcaModel:  onnx.CreateModel("pca"),
		stdModel:  onnx.CreateModel("std"),
		userModel: onnx.CreateModel("user"),
	}
}

func (u *ModelUsecae) CatGroup(input []float64) (int64, error) {
	for i := range input {
		input[i] = (input[i]/10)*(max[i]-min[i]) + min[i]
	}

	var output []float64
	if err := u.stdModel.Run(input, &output); err != nil {
		return 0, err
	}

	if err := u.pcaModel.Run(output, &output); err != nil {
		return 0, err
	}

	var result []int64
	if err := u.catModel.Run(output, &result); err != nil {
		return 0, err
	}

	return result[0], nil
}

func (u *ModelUsecae) UserGroup(input []float64) (int64, error) {
	var result []int64
	if err := u.userModel.Run(input, &result); err != nil {
		return -1, err
	}

	return result[0], nil
}
