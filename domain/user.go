package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/types/date"
)

type UserModel struct {
	ID       int           `json:"id" db:"id"`
	Username string        `json:"username" db:"username"`
	Password string        `json:"password" db:"password"`
	Email    string        `json:"email" db:"email"`
	Salt     string        `json:"salt" db:"salt"`
	Gender   string        `json:"gender" db:"gender"`
	Date     date.JSONDate `json:"date" db:"date"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id int) (*UserModel, error)
	FindByUsername(ctx context.Context, username string) (*UserModel, error)
	FindByEmail(ctx context.Context, email string) (*UserModel, error)
	Insert(ctx context.Context, user *UserModel) (int, error)
	Update(ctx context.Context, user *UserModel) error
	Delete(ctx context.Context, id int) error
}
