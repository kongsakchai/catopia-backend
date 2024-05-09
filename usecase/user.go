package usecase

import (
	"context"
	"fmt"
	"strings"

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

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("user not found"))
	}

	return user, nil
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("user not found"))
	}

	return user, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("user not found"))
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	findEmain, err := u.GetByEmail(ctx, user.Email)
	fmt.Println(findEmain, err)
	if findEmain != nil {
		return errs.NewError(errs.ErrConflict, fmt.Errorf("email already exists"))
	} else if errs.GetCode(err) != errs.ErrNotFound {
		return err
	}

	findUsername, err := u.GetByUsername(ctx, user.Username)
	if findUsername != nil {
		return errs.NewError(errs.ErrConflict, fmt.Errorf("username already exists"))
	} else if errs.GetCode(err) != errs.ErrNotFound {
		return err
	}

	user.Salt = pwd.Salt(15)
	user.Password, err = pwd.PasswordHash(user.Password, user.Salt)
	if err != nil {
		return err
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil

}

func (u *userUsecase) Update(ctx context.Context, id int64, user *domain.User) error {
	find, err := u.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if user.Username != "" {
		find.Username = user.Username
	}

	if user.Email != "" {
		find.Email = user.Email
	}

	if user.Gender != "" {
		find.Gender = user.Gender
	}

	if user.Profile != nil {
		if strings.Compare(*find.Profile, *user.Profile) != 0 {
			u.fileUsecase.RemoveFile(*find.Profile)
		}

		find.Profile = user.Profile
	}

	if !pwd.Compare(user.Password, find.Salt, find.Password) {
		find.Salt = pwd.Salt(15)
		hash, err := pwd.PasswordHash(user.Password, find.Salt)
		if err != nil {
			return err
		}

		find.Password = hash
	}

	return u.userRepo.Update(ctx, find)
}

func (u *userUsecase) ResetPassword(ctx context.Context, code string, password string) error {
	otp, _, err := u.otpUsecase.GetOTP(ctx, code)
	if err != nil {
		return err
	}

	salt := pwd.Salt(15)
	hash, err := pwd.PasswordHash(password, salt)
	if err != nil {
		return err
	}

	err = u.userRepo.UpdatePassword(ctx, otp.ID, hash, salt)
	if err != nil {
		return err
	}
	u.otpUsecase.Delete(ctx, code)

	return nil
}

func (u *userUsecase) ForgetPassword(ctx context.Context, username string) (string, error) {
	user, err := u.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errs.NewError(errs.ErrNotFound, fmt.Errorf("user not found"))
	}

	code, err := u.otpUsecase.Create(ctx, user.ID, user.Email)
	return code, nil
}
