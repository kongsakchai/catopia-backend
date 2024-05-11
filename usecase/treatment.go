package usecase

import (
	"context"
	"fmt"

	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type treatmentUsecase struct {
	treatmentRepo domain.TreatmentRepository
	catUsecase    domain.CatUsecase
}

func NewTreatmentUsecase(t domain.TreatmentRepository, c domain.CatUsecase) domain.TreatmentUsecase {
	return &treatmentUsecase{
		treatmentRepo: t,
		catUsecase:    c,
	}
}

func (t *treatmentUsecase) GetType(ctx context.Context) ([]domain.TreatmentType, error) {
	data, err := t.treatmentRepo.GetType(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *treatmentUsecase) GetByID(ctx context.Context, id int64, catID int64) (*domain.Treatment, error) {
	data, err := t.treatmentRepo.GetByID(ctx, id, catID)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("treatment not found"))
	}

	return data, nil
}

func (t *treatmentUsecase) GetByCatID(ctx context.Context, catID int64, userID int64) ([]domain.Treatment, error) {
	cat, err := t.catUsecase.GetByID(ctx, catID, userID)
	if err != nil {
		return nil, err
	}

	if cat == nil {
		return nil, errs.NewError(errs.ErrNotFound, fmt.Errorf("cat not found"))
	}

	data, err := t.treatmentRepo.GetByCatID(ctx, cat.ID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *treatmentUsecase) Create(ctx context.Context, catID int64, userID int64, treatment *domain.Treatment) error {
	_, err := t.catUsecase.GetByID(ctx, catID, userID)
	if err != nil {
		return err
	}

	treatment.CatID = catID
	err = t.treatmentRepo.Create(ctx, treatment)
	if err != nil {
		return err
	}

	return nil
}

func (t *treatmentUsecase) Update(ctx context.Context, id int64, catID int64, treatment *domain.Treatment) error {
	find, err := t.GetByID(ctx, id, catID)
	if err != nil {
		return err
	}

	if treatment.TreatmentTypeID != find.TreatmentTypeID {
		find.TreatmentTypeID = treatment.TreatmentTypeID
	}

	if treatment.Date.Time().Equal(find.Date.Time()) {
		find.Date = treatment.Date
	}

	if treatment.Location != find.Location {
		find.Location = treatment.Location
	}

	if treatment.Vet != find.Vet {
		find.Vet = treatment.Vet
	}

	if treatment.Detail != find.Detail {
		find.Detail = treatment.Detail
	}

	err = t.treatmentRepo.Update(ctx, find)
	if err != nil {
		return err
	}

	return nil
}

func (t *treatmentUsecase) Delete(ctx context.Context, id int64, catID int64) error {
	_, err := t.GetByID(ctx, id, catID)
	if err != nil {
		return err
	}

	err = t.treatmentRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
