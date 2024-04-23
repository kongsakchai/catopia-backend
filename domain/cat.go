package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type CatModel struct {
	ID         int64          `json:"id" db:"id"`
	Name       string         `json:"name" db:"name"`
	Weight     float64        `json:"weight" db:"weight"`
	Gender     string         `json:"gender" db:"gender"`
	Profile    *string        `json:"profile" db:"profile"`
	Date       *date.JSONDate `json:"date" db:"date"`
	CreateAt   *date.JSONDate `json:"createAt" db:"created_at"`
	UserID     int64          `json:"userId" db:"user_id"`
	LastUpdate *date.JSONDate `json:"last_update" db:"last_update"`
}

type CatRepository interface {
	GetByID(ctx context.Context, id int, userID int) (*CatModel, error)
	GetByUserID(ctx context.Context, userID int) ([]CatModel, error)
	Create(ctx context.Context, cat *CatModel) error
	Update(ctx context.Context, cat *CatModel) error
	Delete(ctx context.Context, id int, userID int) error
}

type CatUsecase interface {
	GetByID(ctx context.Context, id int, userID int) (*CatModel, error)
	GetByUserID(ctx context.Context, userID int) ([]CatModel, error)
	Create(ctx context.Context, cat *CatModel, userID int) error
	Update(ctx context.Context, id int, userID int, cat *CatModel) error
	Delete(ctx context.Context, id int, userID int) error
}
