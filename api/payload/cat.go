package payload

import "github.com/kongsakchai/catopia-backend/domain/date"

type UpdateCat struct {
	Name         string         `json:"name" example:"mori"`
	Gender       string         `json:"gender" example:"male" enums:"male,female"`
	Profile      *string        `json:"profile" example:"url of image"`
	Date         *date.JSONDate `json:"date" example:"2021-01-20" format:"date"`
	Weight       float64        `json:"weight" example:"3.5" format:"float"`
	Breeding     string         `json:"breeding" example:"siamese"`
	Aggression   int            `json:"aggression" example:"5"`
	Shyness      int            `json:"shyness" example:"5"`
	Extraversion int            `json:"extraversion" example:"5"`
}

type CreateCat struct {
	Name         string         `json:"name" binding:"required" example:"mori"`
	Gender       string         `json:"gender" binding:"required" example:"male" enums:"male,female"`
	Profile      *string        `json:"profile" example:"url of image"`
	Date         *date.JSONDate `json:"date" binding:"required" example:"2021-01-20" format:"date"`
	Weight       float64        `json:"weight" binding:"required" example:"3.5" format:"float"`
	Breeding     string         `json:"breeding" binding:"required" example:"siamese"`
	Aggression   int            `json:"aggression" example:"5"`
	Shyness      int            `json:"shyness" example:"5"`
	Extraversion int            `json:"extraversion" example:"5"`
}
