package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type CreateTreatment struct {
	TreatmentTypeID int            `json:"treatmentTypeID" binding:"required" example:"1"`
	Date            *date.JSONDate `json:"date" binding:"required" example:"2021-01-20" format:"date"`
	Location        string         `json:"location" example:"clinic"`
	Vet             string         `json:"vet" example:"Dr. John Doe"`
	Detail          string         `json:"detail" example:"vaccination"`
}

type UpdateTreatment struct {
	TreatmentTypeID int            `json:"treatmentTypeID" example:"1"`
	Date            *date.JSONDate `json:"date" example:"2021-01-20" format:"date"`
	Location        string         `json:"location" example:"clinic"`
	Vet             string         `json:"vet" example:"Dr. John Doe"`
	Detail          string         `json:"detail" example:"vaccination"`
}
