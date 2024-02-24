package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/src/dto"
	"github.com/kongsakchai/catopia-backend/src/helper"
	"github.com/kongsakchai/catopia-backend/src/model"
	"github.com/kongsakchai/catopia-backend/src/usecase"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, helper.BadRequestResponse(err))
	}

	data, err := helper.Mapping[model.UserModel](req)
	if err != nil {
		c.JSON(500, helper.GenerateResponse(false, nil, err.Error(), 500))
		return
	}

	if err := h.authUsecase.SignUp(data); err != nil {
		c.JSON(401, helper.FailSignUpResponse(err))
		return
	}

	c.JSON(200, helper.GenerateResponse(true, nil, "Success", 200))
}
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, helper.BadRequestResponse(err))
		return
	}

	data, err := h.authUsecase.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(401, helper.UnauthorizedResponse(err))
		return
	}

	c.JSON(200, helper.GenerateResponse(true, data, "Success", 200))
}
