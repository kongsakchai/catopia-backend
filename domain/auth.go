package domain

import "context"

type AuthUsecase interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, email, password string) error
	ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error
	ResetPassword(ctx context.Context, username, email string) error
	VerifyEmail(ctx context.Context, username, email string) error
}
