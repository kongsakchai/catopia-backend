package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type catUsecase struct {
	repo        domain.CatRepository
	fileUsecase domain.FileUsecase
	model       domain.ModelUsecae
}

func NewCatUsecase(repo domain.CatRepository, fileUsecase domain.FileUsecase, model domain.ModelUsecae) domain.CatUsecase {
	return &catUsecase{repo, fileUsecase, model}
}

func (u *catUsecase) GetByID(ctx context.Context, id int64, userID int64) (*domain.Cat, error) {
	cat, err := u.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	if cat == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("cat not found"))
	}

	return cat, nil
}

func (u *catUsecase) GetByUserID(ctx context.Context, userID int64) ([]domain.Cat, error) {
	cats, err := u.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (u *catUsecase) Create(ctx context.Context, userID int64, cat *domain.Cat) error {
	cat.UserID = userID
	err := u.repo.Create(ctx, cat)
	if err != nil {
		return err
	}
	return nil
}

func (u *catUsecase) Update(ctx context.Context, id int64, userID int64, cat *domain.Cat) error {
	find, err := u.repo.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	if cat.Profile != nil && find.Profile != nil {
		if strings.Compare(*find.Profile, *cat.Profile) != 0 {
			u.fileUsecase.RemoveFile(*find.Profile)
		}

		find.Profile = cat.Profile
	}

	if cat.Name != "" {
		find.Name = cat.Name
	}

	if cat.Weight != 0 {
		find.Weight = cat.Weight
	}

	if cat.Gender != "" {
		find.Gender = cat.Gender
	}

	if cat.Date != nil {
		find.Date = cat.Date
	}

	if cat.Breeding != "" {
		find.Breeding = cat.Breeding
	}

	find.Aggression = cat.Aggression
	find.Shyness = cat.Shyness
	find.Extraversion = cat.Extraversion

	go func() {
		id, err := u.model.CatGroup([]float64{float64(find.Aggression), float64(find.Shyness), float64(find.Extraversion)})
		if err == nil {
			u.repo.UpdateGroup(ctx, find.ID, id)
		}
	}()

	err = u.repo.Update(ctx, find)
	if err != nil {
		return err
	}

	go func() {
		id, err := u.model.CatGroup([]float64{float64(find.Aggression), float64(find.Shyness), float64(find.Extraversion)})
		if err == nil && id != -1 {
			u.repo.UpdateGroup(ctx, find.ID, id)
		}
	}()

	return nil
}

func (u *catUsecase) Delete(ctx context.Context, id int64, userID int64) error {
	_, err := u.repo.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	err = u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *catUsecase) GetCluster(ctx context.Context, id int64, userID int64) ([]string, error) {
	_, err := u.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	groups, err := u.repo.GetCluster(ctx, id)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (u *catUsecase) GetRandom(ctx context.Context) ([]string, error) {
	groups, err := u.repo.GetRandom(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
