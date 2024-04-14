package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
	secret      string
}

func NewSessionUsecase(sessionRepo domain.SessionRepository) domain.SessionUsecase {
	cfg := config.Get()

	return &sessionUsecase{
		sessionRepo: sessionRepo,
		secret:      cfg.Secret,
	}
}

func (u *sessionUsecase) Signature(id string, expire int64) []byte {
	expireStr := fmt.Sprint(expire)
	hmac := hmac.New(sha256.New, []byte(u.secret))
	hmac.Write([]byte(id))
	hmac.Write([]byte(expireStr))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return []byte(signature)
}

func (u *sessionUsecase) Sign(id string, expire int64) (string, error) {
	signature := u.Signature(id, expire)

	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub": id,
		"exp": expire,
	})

	token, err := payload.SignedString([]byte(signature))
	if err != nil {
		return "", errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return token, nil
}

func (u *sessionUsecase) Create(ctx context.Context, userID int64) (string, error) {
	expireDate := time.Now().Add(30 * time.Minute)
	id := uuid.NewString()

	token, err := u.Sign(id, expireDate.Unix())
	if err != nil {
		return "", err
	}

	session := &domain.Session{
		UserID:    userID,
		Token:     token,
		ID:        id,
		ExpiredAt: date.JSONDate(expireDate),
	}

	err = u.sessionRepo.Create(ctx, session)
	if err != nil {
		return "", errs.New(errs.ErrInternal, "Internal server error", err)
	}

	return token, nil
}

func (u *sessionUsecase) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	session, err := u.sessionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errs.New(errs.ErrInternal, "Internal server error", err)
	}

	if session == nil {
		return nil, errs.New(errs.ErrUnauthorized, "Token not found", nil)
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

func (u *sessionUsecase) UnSign(token string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		exp := int64(claims["exp"].(float64))
		signature := u.Signature(claims["sub"].(string), exp)

		return signature, nil
	})

	if err != nil {
		return "", errs.New(errs.ErrUnauthorized, "Invalid token", err)
	}

	return claims["sub"].(string), nil
}

func (u *sessionUsecase) ValidateToken(ctx context.Context, token string) (*domain.Session, error) {
	id, err := u.UnSign(token)
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

	return session, nil
}
