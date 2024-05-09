package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/utils/mapping"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := h.userUsecase.GetByID(c, id)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req payload.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.User](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = h.userUsecase.Update(c, id, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}

func (h *UserHandler) ForgetPassword(c *gin.Context) {
	var req payload.GetOTP
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	code, err := h.userUsecase.ForgetPassword(c, req.Username)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", gin.H{"code": code})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var req payload.UpdatePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	err := h.userUsecase.ResetPassword(c, req.Code, req.Password)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}

func (h *UserHandler) UserAnswer(c *gin.Context) {
	var req payload.UserAnswer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}
