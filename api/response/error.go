package response

import (
	"net/http"

	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

var ErrCodeToHTTPStatus = map[int]int{
	errs.ErrNotFound:     http.StatusNotFound,
	errs.ErrConflict:     http.StatusConflict,
	errs.ErrBadRequest:   http.StatusBadRequest,
	errs.ErrUnauthorized: http.StatusUnauthorized,

	errs.ErrTransaction: http.StatusInternalServerError,
	errs.ErrRollback:    http.StatusInternalServerError,
	errs.ErrCommit:      http.StatusInternalServerError,

	errs.ErrUserGetByEmail:     http.StatusInternalServerError,
	errs.ErrUserGetByID:        http.StatusInternalServerError,
	errs.ErrUserCreate:         http.StatusInternalServerError,
	errs.ErrUserUpdate:         http.StatusInternalServerError,
	errs.ErrUserGetByUsername:  http.StatusInternalServerError,
	errs.ErrUserUpdatePassword: http.StatusInternalServerError,

	errs.ErrTreatmentTypeGet: http.StatusInternalServerError,
	errs.ErrTreatmentGetByID: http.StatusInternalServerError,
	errs.ErrTreatmentCreate:  http.StatusInternalServerError,
	errs.ErrTreatmentUpdate:  http.StatusInternalServerError,
	errs.ErrTreatmentDelete:  http.StatusInternalServerError,

	errs.ErrSessionCreate: http.StatusInternalServerError,
	errs.ErrSessionGet:    http.StatusInternalServerError,
	errs.ErrSessionDelete: http.StatusInternalServerError,

	errs.ErrCatGetByID:     http.StatusInternalServerError,
	errs.ErrCatGetByUserID: http.StatusInternalServerError,
	errs.ErrCatCreate:      http.StatusInternalServerError,
	errs.ErrCatUpdate:      http.StatusInternalServerError,
	errs.ErrCatDelete:      http.StatusInternalServerError,
	errs.ErrCatUpdateGroup: http.StatusInternalServerError,

	errs.ErrCatGroup: http.StatusInternalServerError,
}
