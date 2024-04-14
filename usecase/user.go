package usecase

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	pwd "github.com/kongsakchai/catopia-backend/utils/password"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo}
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
