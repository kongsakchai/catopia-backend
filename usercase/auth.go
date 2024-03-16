package usercase

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/helper"
	errs "github.com/kongsakchai/catopia-backend/types/error"
)

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (u *authUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", errs.New(500, "internal server error", err)
	}

	if user == nil {
		return "", errs.New(404, "user not found", nil)
	}

	if helper.CheckPasswordHash(password, user.Salt, user.Password) {
		return "", errs.New(401, "invalid username or password", nil)
	}

	return "", nil
}

func (u *authUsecase) Register(ctx context.Context, username, email, password string) error {
	return nil
}

func (u *authUsecase) ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	return nil
}

func (u *authUsecase) ResetPassword(ctx context.Context, username, email string) error {
	return nil
}

func (u *authUsecase) VerifyEmail(ctx context.Context, username, email string) error {
	return nil
}
