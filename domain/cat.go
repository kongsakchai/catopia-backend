package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type Cat struct {
	ID           int64          `json:"id" db:"id"`
	Name         string         `json:"name" db:"name"`
	Weight       float64        `json:"weight" db:"weight"`
	Gender       string         `json:"gender" db:"gender"`
	Profile      *string        `json:"profile" db:"profile"`
	Date         *date.JSONDate `json:"date" db:"date"`
	Breeding     string         `json:"breeding" db:"breeding"`
	Aggression   int            `json:"aggression" db:"aggression"`
	Shyness      int            `json:"shyness" db:"shyness"`
	Extraversion int            `json:"extraversion" db:"extraversion"`
	UserID       int64          `json:"userId" db:"user_id"`
	Group_ID     *int64         `json:"group_id" db:"group_id"`
	LastUpdate   *date.JSONDate `json:"last_update" db:"last_update"`
	CreateAt     *date.JSONDate `json:"createAt" db:"created_at"`
}

type CatRepository interface {
	GetByID(ctx context.Context, id int64, userID int64) (*Cat, error)
	GetByUserID(ctx context.Context, userID int64) ([]Cat, error)
	Create(ctx context.Context, cat *Cat) error
	Update(ctx context.Context, cat *Cat) error
	Delete(ctx context.Context, id int64) error
	UpdateGroup(ctx context.Context, id int64, group int64) error
	GetCluster(ctx context.Context, group int64) ([]string, error)
	GetRandom(ctx context.Context) ([]string, error)
}

type CatUsecase interface {
	GetByID(ctx context.Context, id int64, userID int64) (*Cat, error)
	GetByUserID(ctx context.Context, userID int64) ([]Cat, error)
	Create(ctx context.Context, userID int64, cat *Cat) error
	Update(ctx context.Context, id int64, userID int64, cat *Cat) error
	Delete(ctx context.Context, id int64, userID int64) error
	GetCluster(ctx context.Context, id int64, userID int64) ([]string, error)
	GetRandom(ctx context.Context) ([]string, error)
}
