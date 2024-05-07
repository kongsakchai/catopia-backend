package domain

import "context"

type OTP struct {
	OTP    string `json:"otp"`
	ID     int64  `json:"id"`
	Expire int64  `json:"expire"`
}

type OTPUsecase interface {
	Create(ctx context.Context, id int64, email string) (string, error)
	GetOTP(ctx context.Context, code string) (*OTP, bool, error)
	Delete(ctx context.Context, code string) error
}
