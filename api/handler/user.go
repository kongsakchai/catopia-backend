package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"github.com/kongsakchai/catopia-backend/utils/data"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := h.userUsecase.GetByID(c, id)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req payload.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[domain.UserModel](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = h.userUsecase.Update(c, id, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req payload.UpdatePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	err = h.userUsecase.UpdatePassword(c, id, req.Password)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusCreated, "success", nil)
}
