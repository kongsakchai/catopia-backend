package usecase

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type catUsecase struct {
	repo        domain.CatRepository
	fileUsecase domain.FileUsecase
}

func NewCatUsecase(repo domain.CatRepository, fileUsecase domain.FileUsecase) domain.CatUsecase {
	return &catUsecase{repo, fileUsecase}
}

func (u *catUsecase) GetByID(ctx context.Context, id int, userID int) (*domain.CatModel, error) {
	data, err := u.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if data == nil {
		return nil, errs.New(errs.ErrNotFound, "Cat not found", nil)
	}

	return data, nil
}

func (u *catUsecase) GetByUserID(ctx context.Context, userID int) ([]domain.CatModel, error) {
	data, err := u.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return data, nil
}

func (u *catUsecase) Create(ctx context.Context, cat *domain.CatModel, userID int) error {
	cat.UserID = int64(userID)
	err := u.repo.Create(ctx, cat)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}

func (u *catUsecase) Update(ctx context.Context, id int, userID int, cat *domain.CatModel) error {
	find, err := u.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	if cat.Name != "" && cat.Name != find.Name {
		find.Name = cat.Name
	}

	if cat.Gender != "" && cat.Gender != find.Gender {
		find.Gender = cat.Gender
	}

	if cat.Profile != nil && cat.Profile != find.Profile {
		u.fileUsecase.RemoveFile(*find.Profile)

		find.Profile = cat.Profile
	}

	if !cat.Date.Time().Equal(find.Date.Time()) {
		find.Date = cat.Date
	}

	if cat.Weight != find.Weight {
		find.Weight = cat.Weight
	}

	err = u.repo.Update(ctx, find)
	if err != nil {
		return err
	}

	return nil
}

func (u *catUsecase) Delete(ctx context.Context, id int, userID int) error {
	find, err := u.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	err = u.repo.Delete(ctx, int(find.ID), userID)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}
