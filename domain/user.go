package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type User struct {
	ID        int64          `json:"id" db:"id"`
	Username  string         `json:"username" db:"username"`
	Password  string         `json:"-" db:"password"`
	Email     string         `json:"email" db:"email"`
	Salt      string         `json:"-" db:"salt"`
	Gender    string         `json:"gender" db:"gender"`
	Profile   *string        `json:"profile" db:"profile"`
	GroupID   *int64         `json:"groupID" db:"group_id"`
	Date      *date.JSONDate `json:"date" db:"date"`
	CreatedAt *date.JSONDate `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, id int64, password string, salt string) error
	UpdateGroup(ctx context.Context, id int64, groupID int64) error
	GetUserIDsByGroup(ctx context.Context, groupID int64) ([]int64, error)
}

type UserUsecase interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int64, user *User) error
	ResetPassword(ctx context.Context, code string, password string) error
	ForgetPassword(ctx context.Context, username string) (string, error)
	GetUserIDsInSameGroup(ctx context.Context, id int64) ([]int64, error)
	UpdateGroup(ctx context.Context, id int64, answer []float64) error
	// CreateOTP(ctx context.Context, username string) (string, error)
	// VerifyOTP(ctx context.Context, code string, otp string) error
	// UpdatePasswordWithCode(ctx context.Context, code string, password string) error
}
