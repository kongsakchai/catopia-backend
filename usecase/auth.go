package usecase

import (
	"context"
	"fmt"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	pwd "github.com/kongsakchai/catopia-backend/utils/password"
)

type authUsecase struct {
	userUsecase    domain.UserUsecase
	sessionUsecase domain.SessionUsecase
}

func NewAuthUsecase(userUsecase domain.UserUsecase, sessionUsecase domain.SessionUsecase) domain.AuthUsecase {
	return &authUsecase{userUsecase, sessionUsecase}
}

func (u *authUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.userUsecase.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !pwd.Compare(password, user.Salt, user.Password) {
		return "", errs.New(errs.ErrUnauthorized, "Invalid password", nil)
	}

	token, err := u.sessionUsecase.Create(ctx, user.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *authUsecase) Logout(ctx context.Context, id string) error {
	err := u.sessionUsecase.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *authUsecase) Register(ctx context.Context, user *domain.UserModel) error {
	find, err := u.userUsecase.GetByUsername(ctx, user.Username)
	if err != nil && err.Error() != "User not found" {
		return err
	}

	if find != nil {
		return errs.New(errs.ErrConflict, "Username already exists", nil)
	}

	find, err = u.userUsecase.GetByEmail(ctx, user.Email)
	fmt.Println(err)

	if err != nil && err.Error() != "User not found" {
		return err
	}

	if find != nil {
		return errs.New(errs.ErrConflict, "Email already exists", nil)
	}

	err = u.userUsecase.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
