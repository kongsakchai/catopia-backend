package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type otpUsecase struct {
	otpRepo domain.OTPRepository
}

func NewOTPUsecase(otpRepo domain.OTPRepository) domain.OTPUsecase {
	return &otpUsecase{otpRepo}
}

func (u *otpUsecase) GetByCodeWithExpire(ctx context.Context, code string) (*domain.OTPModel, error) {
	otp, err := u.otpRepo.GetByCode(ctx, code)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "cannot get otp by code", err)
	}

	if otp == nil {
		return nil, errs.New(errs.ErrNotFound, "otp not found", nil)
	}

	if otp.CreatedAt.Time().Add(time.Minute * 3).Before(time.Now()) {
		err = u.Delete(ctx, code)
		return nil, errs.New(errs.ErrNotFound, "OTP expired", nil)
	}

	return otp, nil
}

func (u *otpUsecase) GetByCode(ctx context.Context, code string) (*domain.OTPModel, error) {
	otp, err := u.otpRepo.GetByCode(ctx, code)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "cannot get otp by code", err)
	}

	if otp == nil {
		return nil, errs.New(errs.ErrNotFound, "otp not found", nil)
	}

	return otp, nil
}

func (u *otpUsecase) Create(ctx context.Context, id int) (string, error) {
	uuid := uuid.New().String()
	otp := "1234"

	otpData := &domain.OTPModel{
		Code: uuid,
		OTP:  otp,
		ID:   id,
	}

	err := u.otpRepo.Create(ctx, otpData)
	if err != nil {
		return "", errs.New(errs.ErrInternal, "cannot create otp", err)
	}

	return uuid, nil
}

func (u *otpUsecase) Delete(ctx context.Context, code string) error {
	err := u.otpRepo.Delete(ctx, code)
	if err != nil {
		return errs.New(errs.ErrInternal, "cannot delete otp", err)
	}

	return nil
}
