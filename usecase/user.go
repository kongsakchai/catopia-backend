package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	pwd "github.com/kongsakchai/catopia-backend/utils/password"
)

type userUsecase struct {
	userRepo     domain.UserRepository
	fileUsecase  domain.FileUsecase
	otpUsecase   domain.OTPUsecase
	modelUsecase domain.ModelUsecae
}

func NewUserUsecase(userRepo domain.UserRepository, fileUsecase domain.FileUsecase, otpUsecase domain.OTPUsecase, modelUsecase domain.ModelUsecae) domain.UserUsecase {
	return &userUsecase{userRepo, fileUsecase, otpUsecase, modelUsecase}
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	fmt.Println("user", user, "err", err, err == nil)

	if err != nil {
		fmt.Println("err not null", err)
		return nil, err
	}

	if user == nil {
		fmt.Println("user not found")
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
		if find.Profile != nil && strings.Compare(*find.Profile, *user.Profile) != 0 {
			u.fileUsecase.RemoveFile(*find.Profile)
		}

		find.Profile = user.Profile
	}

	if user.Password != "" && !pwd.Compare(user.Password, find.Salt, find.Password) {
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

	return u.otpUsecase.Create(ctx, user.ID, user.Email)
}

func (u *userUsecase) GetUserIDsInSameGroup(ctx context.Context, id int64) ([]int64, error) {
	user, err := u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.GroupID == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("user not in group"))
	}

	return u.userRepo.GetUserIDsByGroup(ctx, *user.GroupID)
}

func (u *userUsecase) UpdateGroup(ctx context.Context, id int64, answer []float64) error {
	user, err := u.GetByID(ctx, id)
	if err != nil {
		return err
	}

	ages := time.Now().Year() - user.Date.Time().Year()
	if ages <= 15 {
		answer = append([]float64{0}, answer...)
	} else if ages >= 16 && ages <= 20 {
		answer = append([]float64{1}, answer...)
	} else if ages >= 21 && ages <= 25 {
		answer = append([]float64{2}, answer...)
	} else if ages >= 26 && ages <= 30 {
		answer = append([]float64{3}, answer...)
	} else if ages >= 31 && ages <= 40 {
		answer = append([]float64{4}, answer...)
	} else {
		answer = append([]float64{5}, answer...)
	}

	if user.Gender == "male" {
		answer = append([]float64{0}, answer...)
	} else {
		answer = append([]float64{1}, answer...)
	}

	groupID, err := u.modelUsecase.UserGroup(answer)
	if err != nil {
		return err
	}

	return u.userRepo.UpdateGroup(ctx, user.ID, groupID)
}
