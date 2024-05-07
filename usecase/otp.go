package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"gopkg.in/gomail.v2"
)

var list = make(map[string]*domain.OTP)

type otpUsecase struct {
	mail *gomail.Message
}

func NewOTPUsecase() domain.OTPUsecase {
	mail := gomail.NewMessage()
	return &otpUsecase{mail}
}

func (u *otpUsecase) Create(ctx context.Context, id int64, email string) (string, error) {
	uuid := uuid.New().String()
	otp := "1234"

	otpData := &domain.OTP{
		OTP:    otp,
		ID:     id,
		Expire: time.Now().Add(time.Minute * 3).Unix(),
	}

	list[uuid] = otpData
	return uuid, nil
}

func (u *otpUsecase) GetOTP(ctx context.Context, code string) (*domain.OTP, bool, error) {
	otp, ok := list[code]
	if !ok {
		return nil, false, errs.NewError(errs.ErrNotFound, fmt.Errorf("otp not found"))
	}

	if otp.Expire < time.Now().Unix() {
		return otp, false, nil
	}

	return otp, true, nil
}

func (u *otpUsecase) Delete(ctx context.Context, code string) error {
	delete(list, code)
	return nil
}

// func (u *otpUsecase) GetByCodeWithExpire(ctx context.Context, code string) (*domain.OTPModel, error) {
// 	otp, err := u.otpRepo.GetByCode(ctx, code)
// 	if err != nil {
// 		return nil, errs.New(errs.ErrInternal, "cannot get otp by code", err)
// 	}

// 	if otp == nil {
// 		return nil, errs.New(errs.ErrNotFound, "otp not found", nil)
// 	}

// 	if otp.CreatedAt.Time().Add(time.Minute * 3).Before(time.Now()) {
// 		err = u.Delete(ctx, code)
// 		return nil, errs.New(errs.ErrNotFound, "OTP expired", nil)
// 	}

// 	return otp, nil
// }

// func (u *otpUsecase) GetByCode(ctx context.Context, code string) (*domain.OTPModel, error) {
// 	otp, err := u.otpRepo.GetByCode(ctx, code)
// 	if err != nil {
// 		return nil, errs.New(errs.ErrInternal, "cannot get otp by code", err)
// 	}

// 	if otp == nil {
// 		return nil, errs.New(errs.ErrNotFound, "otp not found", nil)
// 	}

// 	return otp, nil
// }

// func (u *otpUsecase) Create(ctx context.Context, id int, email string) (string, error) {
// 	uuid := uuid.New().String()
// 	otp := "1234"

// 	otpData := &domain.OTP{
// 		OTP:    otp,
// 		ID:     id,
// 		Expire: time.Now().Add(time.Minute * 3).Unix(),
// 	}

// 	return uuid, nil
// }

// func (u *otpUsecase) Delete(ctx context.Context, code string) error {
// 	err := u.otpRepo.Delete(ctx, code)
// 	if err != nil {
// 		return errs.New(errs.ErrInternal, "cannot delete otp", err)
// 	}

// 	return nil
// }
