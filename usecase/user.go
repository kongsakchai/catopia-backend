package usecase

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	pwd "github.com/kongsakchai/catopia-backend/utils/password"
)

type userUsecase struct {
	userRepo    domain.UserRepository
	fileUsecase domain.FileUsecase
	otpUsecase  domain.OTPUsecase
}

func NewUserUsecase(userRepo domain.UserRepository, fileUsecase domain.FileUsecase, otpUsecase domain.OTPUsecase) domain.UserUsecase {
	return &userUsecase{userRepo, fileUsecase, otpUsecase}
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	data, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if data == nil {
		return nil, errs.New(errs.ErrNotFound, "User not found", nil)
	}

	return data, nil
}
func (u *userUsecase) GetByID(ctx context.Context, id int) (*domain.UserModel, error) {
	data, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if data == nil {
		return nil, errs.New(errs.ErrNotFound, "User not found", nil)
	}

	return data, nil
}

func (u *userUsecase) Create(ctx context.Context, user *domain.UserModel) error {
	salt := pwd.Salt(9)
	hashPassword, err := pwd.PasswordHash(user.Password, salt)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	user.Password = hashPassword
	user.Salt = salt

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}

func (u *userUsecase) Update(ctx context.Context, id int, data *domain.UserModel) error {
	user, err := u.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if data.Password != "" {
		salt := pwd.Salt(9)
		hashPassword, err := pwd.PasswordHash(user.Password, salt)
		if err != nil {
			return errs.New(errs.ErrInternal, "Internal server error", err)
		}

		user.Password = hashPassword
		user.Salt = salt
	}

	if data.Username != "" && data.Username != user.Username {
		user.Username = data.Username
	}

	if data.Email != "" && data.Email != user.Email {
		user.Email = data.Email
	}

	if data.Gender != "" && data.Gender != user.Gender {
		user.Gender = data.Gender
	}

	if !data.Date.Time().Equal(user.Date.Time()) {
		user.Date = data.Date
	}

	if data.Profile != nil && data.Profile != user.Profile {
		if user.Profile != nil {
			u.fileUsecase.RemoveFile(*user.Profile)
		}

		user.Profile = data.Profile
	}

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}
	return nil
}

func (u *userUsecase) UpdatePassword(ctx context.Context, id int, password string) error {
	data, err := u.GetByID(ctx, id)
	if err != nil {
		return err
	}

	salt := pwd.Salt(9)
	hashPassword, err := pwd.PasswordHash(password, salt)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	data.Password = hashPassword
	data.Salt = salt

	err = u.userRepo.Update(ctx, data)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}
	return nil
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (*domain.UserModel, error) {
	data, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if data == nil {
		return nil, errs.New(errs.ErrNotFound, "User not found", nil)
	}

	return data, nil
}

func (u *userUsecase) CreateOTP(ctx context.Context, username string) (string, error) {
	data, err := u.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	code, err := u.otpUsecase.Create(ctx, int(data.ID))
	if err != nil {
		return "", err
	}

	return code, nil
}

func (u *userUsecase) VerifyOTP(ctx context.Context, code string, otp string) error {
	data, err := u.otpUsecase.GetByCodeWithExpire(ctx, code)
	if err != nil {
		return err
	}

	if data.OTP != otp {
		return errs.New(errs.ErrNotFound, "OTP not match", nil)
	}

	return nil
}

func (u *userUsecase) UpdatePasswordWithCode(ctx context.Context, code string, password string) error {
	data, err := u.otpUsecase.GetByCode(ctx, code)
	if err != nil {
		return err
	}
	u.otpUsecase.Delete(ctx, code)

	return u.UpdatePassword(ctx, data.ID, password)
}
