package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type TreatmentType struct {
	ID            int64  `json:"id" db:"id"`
	TreatmentType string `json:"treatment_type" db:"treatment_type"`
}

type Treatment struct {
	ID              int64              `json:"id" db:"id"`
	CatID           int64              `json:"catID" db:"cat_id"`
	Name            string             `json:"name" db:"name"`
	TreatmentTypeID int64              `json:"treatmentTypeID" db:"treatment_type_id"`
	Date            *date.JSONDate     `json:"date" db:"date"`
	Location        string             `json:"location" db:"location"`
	Vet             string             `json:"vet" db:"vet"`
	Detail          string             `json:"detail" db:"detail"`
	Appointment     *string            `json:"appointment" db:"appointment"`
	AppointmentDate *date.JSONDateTime `json:"appointmentDate" db:"appointment_date"`
	CreatedAt       *date.JSONDate     `json:"createdAt" db:"created_at"`
}

type TreatmentNoti struct {
	ID              int64              `json:"id" db:"id"`
	CatID           int64              `json:"catID" db:"cat_id"`
	Name            string             `json:"name" db:"name"`
	Appointment     *string            `json:"appointment" db:"appointment"`
	AppointmentDate *date.JSONDateTime `json:"appointmentDate" db:"appointment_date"`
}

type TreatmentRepository interface {
	GetType(ctx context.Context) ([]TreatmentType, error)
	GetByID(ctx context.Context, id int64, catID int64) (*Treatment, error)
	GetByCatID(ctx context.Context, catID int64) ([]Treatment, error)
	Create(ctx context.Context, treatment *Treatment) error
	Update(ctx context.Context, treatment *Treatment) error
	Delete(ctx context.Context, id int64) error
	GetTreatmentNoti(ctx context.Context, userID int64) ([]TreatmentNoti, error)
}

type TreatmentUsecase interface {
	GetType(ctx context.Context) ([]TreatmentType, error)
	GetByID(ctx context.Context, id int64, catID int64) (*Treatment, error)
	GetByCatID(ctx context.Context, catID int64, userID int64) ([]Treatment, error)
	Create(ctx context.Context, catID int64, userID int64, treatment *Treatment) error
	Update(ctx context.Context, id int64, catID int64, treatment *Treatment) error
	Delete(ctx context.Context, id int64, catID int64) error
	GetTreatmentNoti(ctx context.Context, userID int64) ([]TreatmentNoti, error)
}
