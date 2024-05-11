package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type catUsecase struct {
	catRepo      domain.CatRepository
	fileUsecase  domain.FileUsecase
	modelUsecase domain.ModelUsecae
}

func NewCatUsecase(catRepo domain.CatRepository, fileUsecase domain.FileUsecase, modelUsecase domain.ModelUsecae) domain.CatUsecase {
	return &catUsecase{catRepo, fileUsecase, modelUsecase}
}

func (u *catUsecase) GetByID(ctx context.Context, id int64, userID int64) (*domain.Cat, error) {
	cat, err := u.catRepo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	if cat == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("cat not found"))
	}

	return cat, nil
}

func (u *catUsecase) GetByUserID(ctx context.Context, userID int64) ([]domain.Cat, error) {
	cats, err := u.catRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (u *catUsecase) Create(ctx context.Context, userID int64, cat *domain.Cat) error {
	cat.UserID = userID
	err := u.catRepo.Create(ctx, cat)
	if err != nil {
		return err
	}

	go func() {
		id, err := u.modelUsecase.CatGroup([]float64{float64(cat.Aggression), float64(cat.Shyness), float64(cat.Extraversion)})
		if err == nil {
			u.catRepo.UpdateGroup(ctx, cat.ID, id)
		}
	}()

	return nil
}

func (u *catUsecase) Update(ctx context.Context, id int64, userID int64, cat *domain.Cat) error {
	find, err := u.catRepo.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	if cat.Profile != nil {
		if find.Profile != nil && strings.Compare(*find.Profile, *cat.Profile) != 0 {
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
		id, err := u.modelUsecase.CatGroup([]float64{float64(find.Aggression), float64(find.Shyness), float64(find.Extraversion)})
		if err == nil {
			u.catRepo.UpdateGroup(ctx, find.ID, id)
		}
	}()

	err = u.catRepo.Update(ctx, find)
	if err != nil {
		return err
	}

	return nil
}

func (u *catUsecase) Delete(ctx context.Context, id int64, userID int64) error {
	_, err := u.catRepo.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	err = u.catRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *catUsecase) GetBreedingByCat(ctx context.Context, id int64, userID int64) ([]string, error) {
	find, err := u.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	var groupID int64
	if find.Group_ID == nil {
		groupID = 0
	} else {
		groupID = *find.Group_ID
	}

	groups, err := u.catRepo.GetBreedingByGroup(ctx, []int64{groupID}, false)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (u *catUsecase) GetBreedingByUser(ctx context.Context, userID []int64) ([]string, error) {
	cats, err := u.catRepo.GetByUserIDs(ctx, userID)
	if err != nil {
		return nil, err
	}

	ids := []int64{}
	for _, cat := range cats {
		if cat.Group_ID != nil {
			ids = append(ids, *cat.Group_ID)
		}
	}

	groups, err := u.catRepo.GetBreedingByGroup(ctx, ids, true)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (u *catUsecase) GetBreedingByRandom(ctx context.Context) ([]string, error) {
	groups, err := u.catRepo.GetBreedingByRandom(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
