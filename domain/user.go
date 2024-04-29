package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type UserModel struct {
	ID        int64          `json:"id" db:"id"`
	Username  string         `json:"username" db:"username"`
	Password  string         `json:"-" db:"password"`
	Email     string         `json:"email" db:"email"`
	Salt      string         `json:"-" db:"salt"`
	Gender    string         `json:"gender" db:"gender"`
	Date      *date.JSONDate `json:"date" db:"date"`
	CreatedAt *date.JSONDate `json:"createdAt" db:"created_at"`
	Profile   *string        `json:"profile" db:"profile"`
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*UserModel, error)
	GetByUsername(ctx context.Context, username string) (*UserModel, error)
	GetByID(ctx context.Context, id int) (*UserModel, error)
	Create(ctx context.Context, user *UserModel) error
	Update(ctx context.Context, user *UserModel) error
}

type UserUsecase interface {
	GetByEmail(ctx context.Context, email string) (*UserModel, error)
	GetByUsername(ctx context.Context, username string) (*UserModel, error)
	GetByID(ctx context.Context, id int) (*UserModel, error)
	Create(ctx context.Context, user *UserModel) error
	Update(ctx context.Context, id int, user *UserModel) error
	UpdatePassword(ctx context.Context, id int, password string) error
	CreateOTP(ctx context.Context, username string) (string, error)
	VerifyOTP(ctx context.Context, code string, otp string) error
	UpdatePasswordWithCode(ctx context.Context, code string, password string) error
}
