package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type CreateTreatment struct {
	TreatmentTypeID int            `json:"treatmentTypeID" binding:"required"`
	Date            *date.JSONDate `json:"date" binding:"required"`
	Location        string         `json:"location" `
	Vet             string         `json:"vet"`
	Detail          string         `json:"detail"`
}

type UpdateTreatment struct {
	TreatmentTypeID int            `json:"treatmentTypeID"`
	Date            *date.JSONDate `json:"date"`
	Location        string         `json:"location"`
	Vet             string         `json:"vet"`
	Detail          string         `json:"detail"`
}
