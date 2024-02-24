package helper

import (
	"fmt"

	"github.com/kongsakchai/catopia-backend/src/dto"
)

func GenerateResponse(success bool, data interface{}, message string, code int) *dto.Response {
	return &dto.Response{
		Sucess:  success,
		Result:  data,
		Message: message,
		Code:    code,
	}
}

func InternalServerErrorResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return GenerateResponse(false, nil, "Internal Server Error", 500)
}

func BadRequestResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return GenerateResponse(false, nil, "Bad Request", 400)
}

func NotFoundResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return GenerateResponse(false, nil, "Not Found", 404)
}

func UnauthorizedResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return GenerateResponse(false, nil, "Unauthorized", 401)
}

func FailSignUpResponse(err error) *dto.Response {
	if err != nil {
		fmt.Println("\u001b[31mERROR: \u001b[0m", err)
	}

	return GenerateResponse(false, nil, err.Error(), 401)
}
