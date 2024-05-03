package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateCat struct {
	Name         string         `json:"name"`
	Gender       string         `json:"gender"`
	Profile      string         `json:"profile"`
	Date         *date.JSONDate `json:"date"`
	Weight       float64        `json:"weight"`
	Breeding     string         `json:"breeding"`
	Aggression   int            `json:"aggression"`
	Shyness      int            `json:"shyness"`
	Extraversion int            `json:"extraversion"`
}

type CreateCat struct {
	Name         string         `json:"name" binding:"required"`
	Gender       string         `json:"gender" binding:"required"`
	Profile      string         `json:"profile" `
	Date         *date.JSONDate `json:"date" binding:"required"`
	Weight       float64        `json:"weight" binding:"required"`
	Breeding     string         `json:"breeding"`
	Aggression   int            `json:"aggression"`
	Shyness      int            `json:"shyness"`
	Extraversion int            `json:"extraversion"`
}
