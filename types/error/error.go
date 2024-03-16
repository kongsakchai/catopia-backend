package errs

import "log"

type Errs struct {
	code    int
	message string
}

func New(code int, message string, err error) *Errs {
	if err != nil {
		log.Println(err)
	}

	return &Errs{code: code, message: message}
}

func (e *Errs) Error() string {
	return e.message
}
