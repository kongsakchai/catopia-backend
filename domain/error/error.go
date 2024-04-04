package errs

import "log"

type Error struct {
	Code    int
	Message string
}

func New(code int, message string, err error) *Error {
	if err != nil {
		log.Println(err.Error())
	}

	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
