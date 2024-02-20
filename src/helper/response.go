package helper

import (
	"fmt"

	"github.com/kongsakchai/catopia-backend/src/dto"
)

func GenerateResponse(success bool, data interface{}, message string) *dto.Response {
	return &dto.Response{
		Sucess:  success,
		Result:  data,
		Message: message,
	}
}

func InternalServerErrorResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return &dto.Response{
		Sucess:  false,
		Result:  nil,
		Message: "Internal Server Error",
	}
}

func BadRequestResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return &dto.Response{
		Sucess:  false,
		Result:  nil,
		Message: "Bad Request",
	}
}

func NotFoundResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return &dto.Response{
		Sucess:  false,
		Result:  nil,
		Message: "Not Found",
	}
}
