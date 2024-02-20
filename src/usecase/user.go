package usecase

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/src/dto"
	"github.com/kongsakchai/catopia-backend/src/helper"
	"github.com/kongsakchai/catopia-backend/src/model"
	"github.com/kongsakchai/catopia-backend/src/repository"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) SignUp(c *gin.Context, req *dto.SignUpDTO) error {
	findUser, err := u.repo.GetByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if findUser.Username != "" {
		return nil
	}

	findUser, err = u.repo.GetByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if findUser.Username != "" {
		return nil
	}

	println(findUser)

	user, err := helper.Mapping[model.UserModel](req)
	if err != nil {
		return err
	}

	user.Salt = helper.RandSalt(9)
	user.Password, err = helper.PasswordHash(user.Password, user.Salt)
	if err != nil {
		return err
	}
	return u.repo.Create(user)

}

func (u *UserUsecase) SignIn(c *gin.Context, req *dto.SignInDTO) (*dto.UserDTO, error) {
	user, err := u.repo.GetByUsername(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if helper.CheckPasswordHash(req.Password+user.Salt, user.Password) {
		return helper.Mapping[dto.UserDTO](user)
	}

	return nil, nil
}
