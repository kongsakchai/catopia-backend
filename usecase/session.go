package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kongsakchai/catopia-backend/config"
	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/domain/date"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type sessionUsecase struct {
	sessionRepo domain.SessionRepository
}

func NewSessionUsecase(sessionRepo domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{sessionRepo}
}

func (u *sessionUsecase) getSecret() []byte {
	cfg := config.Get()
	hmac := hmac.New(sha256.New, []byte(cfg.Secret))
	return hmac.Sum(nil)
}

func (u *sessionUsecase) Create(ctx context.Context, userID int64) (string, error) {
	expireDate := time.Now().Add(30 * time.Minute)
	secret := u.getSecret()
	id := uuid.NewString()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id": id,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errs.New(errs.ErrInternal, "Internal server error", err)
	}

	session := &domain.Session{
		UserID:    userID,
		Token:     tokenString,
		ID:        id,
		ExpiredAt: date.JSONDate(expireDate),
	}

	err = u.sessionRepo.Create(ctx, session)
	if err != nil {
		return "", errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return tokenString, nil
}

func (u *sessionUsecase) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	session, err := u.sessionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if session == nil {
		return nil, errs.New(errs.ErrNotFound, "Token not found", nil)
	}

	return session, nil

}

func (u *sessionUsecase) Delete(ctx context.Context, id string) error {
	err := u.sessionRepo.Delete(ctx, id)
	if err != nil {
		return errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return nil
}

func (u *sessionUsecase) ParseToken(ctx context.Context, token string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return u.getSecret(), nil
	})

	if err != nil {
		return "", errs.New(errs.ErrUnauthorized, "Invalid token", err)
	}

	return claims["id"].(string), nil
}

func (u *sessionUsecase) ValidateToken(ctx context.Context, token string) (*domain.Session, error) {
	id, err := u.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}

	session, err := u.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if time.Now().After(time.Time(session.ExpiredAt)) {
		u.Delete(ctx, session.ID)

		return nil, errs.New(errs.ErrUnauthorized, "Token expired", nil)
	}

	// if len(session.Token) != len(token) && strings.Compare(session.Token, token) != 0 {
	// 	u.Delete(ctx, session.ID)

	// 	return nil, errs.New(errs.ErrUnauthorized, "Unauthorized", nil)
	// }

	return session, nil
}
