package domain

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

type Session struct {
	ID        string         `json:"id" db:"id"`
	UserID    int64          `json:"userId" db:"user_id"`
	Token     string         `json:"userAgent" db:"token"`
	ExpiredAt date.JSONDate  `json:"expiredAt" db:"expired_at"`
	CreatedAt *date.JSONDate `json:"createdAt" db:"created_at"`
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByID(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
	ClearExpired(ctx context.Context) error
}

type SessionUsecase interface {
	Create(ctx context.Context, userID int64) (string, error)
	FindByID(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
	ValidateToken(ctx context.Context, token string) (*Session, error)
}
