package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type OTPModel struct {
	Code      string         `json:"code" db:"code"`
	OTP       string         `json:"otp" db:"otp"`
	ID        int            `json:"id" db:"id"`
	CreatedAt *date.JSONDate `json:"created_at" db:"created_at"`
}

type OTPRepository interface {
	GetByCode(ctx context.Context, string string) (*OTPModel, error)
	Create(ctx context.Context, otp *OTPModel) error
	Delete(ctx context.Context, code string) error
}

type OTPUsecase interface {
	GetByCodeWithExpire(ctx context.Context, code string) (*OTPModel, error)
	GetByCode(ctx context.Context, code string) (*OTPModel, error)
	Create(ctx context.Context, id int) (string, error)
	Delete(ctx context.Context, code string) error
}
