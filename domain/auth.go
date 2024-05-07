package domain

import "context"

type AuthUsecase interface {
	Login(ctx context.Context, username, password string) (string, error)
	Logout(ctx context.Context, id string) error
	Register(ctx context.Context, user *User) error
}
