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

// Get godoc
// @description Get user detail
// @tags user
// @security ApiKeyAuth
// @id UserGetHandler
// @accept json
// @produce json
// @Router /api/user [get]
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

// Update godoc
// @description Update user detail
// @tags user
// @security ApiKeyAuth
// @id UserUpdateHandler
// @accept json
// @produce json
// @param user_id path int true "user id"
// @param body body payload.UpdateUser true "user data"
// @Router /api/user [put]
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

// ForgetPassword godoc
// @description Forget password
// @tags Forgot Password
// @id UserForgetPasswordHandler
// @accept json
// @produce json
// @param body body payload.GetOTP true "username"
// @Router /api/forget-password [post]
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

// ResetPassword godoc
// @description Reset password
// @tags Forgot Password
// @id UserResetPasswordHandler
// @accept json
// @produce json
// @param body body payload.UpdatePassword true "code and password"
// @Router /api/reset-password [put]
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

// UserAnswer godoc
// @description User answer
// @tags user
// @security ApiKeyAuth
// @id UserAnswerHandler
// @accept json
// @produce json
// @param body body payload.UserAnswer true "user answer"
// @Router /api/user/answer [post]
func (h *UserHandler) UserAnswer(c *gin.Context) {
	var req payload.UserAnswer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = h.userUsecase.UpdateGroup(c, id, req.Answer)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}
