package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type TreatmentTypeModel struct {
	ID            int    `json:"id" db:"id"`
	TreatmentType string `json:"treatment_type" db:"treatment_type"`
	CreatedAt     string `json:"created_at" db:"created_at"`
}

type TreatmentModel struct {
	ID              int            `json:"id" db:"id"`
	CatID           int            `json:"catID" db:"cat_id"`
	TreatmentTypeID int            `json:"treatmentTypeID" db:"treatment_type_id"`
	Weight          float64        `json:"weight" db:"weight"`
	Date            *date.JSONDate `json:"date" db:"date"`
	Location        string         `json:"location" db:"location"`
	Vet             string         `json:"vet" db:"vet"`
	Detail          string         `json:"detail" db:"detail"`
	CreatedAt       *date.JSONDate `json:"createdAt" db:"created_at"`
}

type TreatmentRepository interface {
	GetType(ctx context.Context) ([]TreatmentTypeModel, error)
	GetByID(ctx context.Context, id int, catID int) (*TreatmentModel, error)
	GetByCatID(ctx context.Context, catID int) ([]TreatmentModel, error)
	Create(ctx context.Context, treatment *TreatmentModel) error
	Update(ctx context.Context, treatment *TreatmentModel) error
	Delete(ctx context.Context, id int, catID int) error
}

type TreatmentUsecase interface {
	GetType(ctx context.Context) ([]TreatmentTypeModel, error)
	GetByID(ctx context.Context, id int, catID int) (*TreatmentModel, error)
	GetByCatID(ctx context.Context, catID int, userID int) ([]TreatmentModel, error)
	Create(ctx context.Context, catID int, userID int, treatment *TreatmentModel) error
	Update(ctx context.Context, id int, catID int, treatment *TreatmentModel) error
	Delete(ctx context.Context, id int, catID int) error
}
