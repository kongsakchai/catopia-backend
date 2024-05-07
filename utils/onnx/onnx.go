package onnx

import (
	"github.com/kongsakchai/catopia-backend/config"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	ort "github.com/yalue/onnxruntime_go"
)

type Model interface {
	Run(input []float64, output interface{}) error
}

type Onnx[Output ort.TensorData] struct {
	Name        string
	ModelPath   string
	InputName   string
	OutputName  string
	InputShape  ort.Shape
	OutputShape ort.Shape
}

func CreateModel(model string) Model {
	if model == "cat" {
		return newCatOnnx()
	}

	if model == "std" {
		return newStdOnnx()
	}

	if model == "pca" {
		return newPCAOnnx()
	}

	if model == "user" {
		return newUserOnnx()
	}

	return nil
}

func newCatOnnx() *Onnx[int64] {
	cfg := config.Get()

	return &Onnx[int64]{
		Name:        "cat",
		ModelPath:   cfg.ModelPath + "cat.onnx",
		InputName:   "X",
		OutputName:  "label",
		InputShape:  ort.NewShape(1, 2),
		OutputShape: ort.NewShape(1),
	}
}

func newStdOnnx() *Onnx[float64] {
	cfg := config.Get()

	return &Onnx[float64]{
		Name:        "std",
		ModelPath:   cfg.ModelPath + "std.onnx",
		InputName:   "X",
		OutputName:  "variable",
		InputShape:  ort.NewShape(1, 3),
		OutputShape: ort.NewShape(1, 3),
	}
}

func newPCAOnnx() *Onnx[float64] {
	cfg := config.Get()

	return &Onnx[float64]{
		Name:        "pca",
		ModelPath:   cfg.ModelPath + "pca.onnx",
		InputName:   "X",
		OutputName:  "variable",
		InputShape:  ort.NewShape(1, 3),
		OutputShape: ort.NewShape(1, 2),
	}
}

func newUserOnnx() *Onnx[int64] {
	cfg := config.Get()

	return &Onnx[int64]{
		Name:        "user",
		ModelPath:   cfg.ModelPath + "user.onnx",
		InputName:   "X",
		OutputName:  "label",
		InputShape:  ort.NewShape(1, 10),
		OutputShape: ort.NewShape(1, 2),
	}
}

func (o *Onnx[Output]) Run(input []float64, output interface{}) error {
	inputTensor, err := ort.NewTensor(o.InputShape, input)
	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}
	defer inputTensor.Destroy()

	outputTensor, err := ort.NewEmptyTensor[Output](o.OutputShape)
	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}
	defer outputTensor.Destroy()

	session, err := ort.NewAdvancedSession(o.ModelPath,
		[]string{o.InputName}, []string{o.OutputName},
		[]ort.ArbitraryTensor{inputTensor}, []ort.ArbitraryTensor{outputTensor}, nil)
	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}
	defer session.Destroy()

	err = session.Run()
	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}

	*output.(*[]Output) = outputTensor.GetData()
	return nil
}
