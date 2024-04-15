package usecase

import (
	"context"

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

func (t *treatmentUsecase) GetType(ctx context.Context) ([]domain.TreatmentTypeModel, error) {
	data, err := t.treatmentRepo.GetType(ctx)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return data, nil
}

func (t *treatmentUsecase) GetByID(ctx context.Context, id int, catID int) (*domain.TreatmentModel, error) {
	data, err := t.treatmentRepo.GetByID(ctx, id, catID)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if data == nil {
		return nil, errs.New(errs.ErrNotFound, "Treatment not found", nil)
	}

	return data, nil
}

func (t *treatmentUsecase) GetByCatID(ctx context.Context, catID int, userID int) ([]domain.TreatmentModel, error) {
	cat, err := t.catUsecase.GetByID(ctx, catID, userID)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if cat == nil {
		return nil, errs.New(errs.ErrNotFound, "Cat not found", nil)
	}

	data, err := t.treatmentRepo.GetByCatID(ctx, int(cat.ID))
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return data, nil
}

func (t *treatmentUsecase) Create(ctx context.Context, catID int, userID int, treatment *domain.TreatmentModel) error {
	cat, err := t.catUsecase.GetByID(ctx, catID, userID)
	if err != nil {
		return err
	}

	treatment.CatID = int(cat.ID)
	err = t.treatmentRepo.Create(ctx, treatment)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}

func (t *treatmentUsecase) Update(ctx context.Context, id int, catID int, treatment *domain.TreatmentModel) error {
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
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}

func (t *treatmentUsecase) Delete(ctx context.Context, id int, catID int) error {
	treatment, err := t.GetByID(ctx, id, catID)
	if err != nil {
		return err
	}

	err = t.treatmentRepo.Delete(ctx, id, int(treatment.CatID))
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}
