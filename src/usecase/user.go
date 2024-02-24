package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/src/model"
	"github.com/kongsakchai/catopia-backend/src/repository"
)

type UserUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(c *gin.Context, userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo}
}

func (u *UserUsecase) Create(user *model.UserModel) error {
	return u.userRepo.Create(user)
}

func (u *UserUsecase) GetByUsername(username string) (*model.UserModel, error) {
	return u.userRepo.GetByUsername(username)
}

func (u *UserUsecase) GetByEmail(email string) (*model.UserModel, error) {
	return u.userRepo.GetByEmail(email)
}

func (u *UserUsecase) GetByID(id int) (*model.UserModel, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserUsecase) Update(user *model.UserModel) error {
	return u.userRepo.Update(user)
}

func (u *UserUsecase) Delete(id int) error {
	return u.userRepo.Delete(id)
}
