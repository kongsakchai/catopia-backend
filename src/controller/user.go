package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/src/usecase"
)

type UserContoller struct {
	userUsecase *usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserContoller {
	return &UserContoller{userUsecase}
}

func (u *UserContoller) SignUp(c *gin.Context) {
	Create(c, u.userUsecase.SignUp)
}

func (u *UserContoller) SignIn(c *gin.Context) {
	GetByBody(c, u.userUsecase.SignIn)
}
