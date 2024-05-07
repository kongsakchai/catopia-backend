package errs

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewError(code int, err error) *Error {
	var message string

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		upwrapErr, ok := err.(*Error)
		if ok {
			return upwrapErr
		}

		message = err.Error()

		_, fn, line, _ := runtime.Caller(1)
		log.Printf("Error at %s:%d: %s \n", fn, line, message)
	}

	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewErrorWithSkip(code int, err error, skip int) *Error {
	var message string

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		upwrapErr, ok := err.(*Error)
		if ok {
			return upwrapErr
		}

		message = err.Error()

		_, fn, line, _ := runtime.Caller(skip)
		log.Printf("Error at %s:%d: %s \n", fn, line, message)
	}

	return &Error{
		Code:    code,
		Message: message,
	}
}

func GetCode(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return 0
	}

	return e.Code
}
