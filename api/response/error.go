package response

import (
	"net/http"

	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

var ErrCodeToHTTPStatus = map[int]int{
	errs.ErrNotFound:   http.StatusNotFound,
	errs.ErrConflict:   http.StatusConflict,
	errs.ErrBadRequest: http.StatusBadRequest,

	errs.ErrTransaction: http.StatusBadRequest,
	errs.ErrRollback:    http.StatusBadRequest,
	errs.ErrCommit:      http.StatusBadRequest,

	errs.ErrUserGetByEmail:     http.StatusBadRequest,
	errs.ErrUserGetByID:        http.StatusBadRequest,
	errs.ErrUserCreate:         http.StatusBadRequest,
	errs.ErrUserUpdate:         http.StatusBadRequest,
	errs.ErrUserGetByUsername:  http.StatusBadRequest,
	errs.ErrUserUpdatePassword: http.StatusBadRequest,

	errs.ErrTreatmentTypeGet: http.StatusBadRequest,
	errs.ErrTreatmentGetByID: http.StatusBadRequest,
	errs.ErrTreatmentCreate:  http.StatusBadRequest,
	errs.ErrTreatmentUpdate:  http.StatusBadRequest,
	errs.ErrTreatmentDelete:  http.StatusBadRequest,

	errs.ErrSessionCreate: http.StatusBadRequest,
	errs.ErrSessionGet:    http.StatusBadRequest,
	errs.ErrSessionDelete: http.StatusBadRequest,

	errs.ErrCatGetByID:     http.StatusBadRequest,
	errs.ErrCatGetByUserID: http.StatusBadRequest,
	errs.ErrCatCreate:      http.StatusBadRequest,
	errs.ErrCatUpdate:      http.StatusBadRequest,
	errs.ErrCatDelete:      http.StatusBadRequest,
	errs.ErrCatUpdateGroup: http.StatusBadRequest,

	errs.ErrCatGroup: http.StatusBadRequest,
}
