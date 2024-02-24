package usecase

import (
	"database/sql"
	"errors"

	"github.com/kongsakchai/catopia-backend/src/helper"
	"github.com/kongsakchai/catopia-backend/src/model"
	"github.com/kongsakchai/catopia-backend/src/repository"
)

type AuthUsecase struct {
	userRepo *repository.UserRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo}
}

func (u *AuthUsecase) SignUp(user *model.UserModel) error {
	findUser, err := u.userRepo.GetByUsername(user.Username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if findUser != nil {
		return errors.New("username already exists")
	}

	findUser, err = u.userRepo.GetByEmail(user.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if findUser != nil {
		return errors.New("email already exists")
	}

	user.Salt = helper.RandSalt(9)
	user.Password, err = helper.PasswordHash(user.Password, user.Salt)
	if err != nil {
		return err
	}

	return u.userRepo.Create(user)
}

func (u *AuthUsecase) SignIn(email string, password string) (*model.UserModel, error) {
	user, err := u.userRepo.GetByUsername(email)
	if err != nil {
		return nil, err
	}

	if helper.CheckPasswordHash(password+user.Salt, user.Password) {
		return user, nil
	}

	return nil, errors.New("invalid email or password")
}
