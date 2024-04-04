package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/response"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"github.com/kongsakchai/catopia-backend/utils/data"
)

func Create[Req any, Data any](c *gin.Context, caller func(c context.Context, data *Data) error) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[Data](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = caller(c, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusCreated, "success", nil)
}

func Update[Req any, Data any](c *gin.Context, caller func(c context.Context, id int, data *Data) error) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[Data](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = caller(c, id, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}

func GetByID[Res any](c *gin.Context, caller func(c context.Context, id int) (*Res, error)) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := caller(c, id)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}
