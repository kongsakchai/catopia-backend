package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
)

type otpRepository struct {
	db *db.Database
}

func NewOTPRepository() domain.OTPRepository {
	db := db.GetDB()
	return &otpRepository{db}
}

func (r *otpRepository) GetByCode(ctx context.Context, code string) (*domain.OTPModel, error) {
	sqlBuild := sq.Select("*").From("otp").Where(sq.Eq{"code": code})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get otp by code:  cannot build sql: %w", err)
	}

	otp := &domain.OTPModel{}
	err = r.db.GetContext(ctx, otp, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get otp by code:  cannot execute query: %w", err)
	}

	return otp, nil
}

func (r *otpRepository) Create(ctx context.Context, otp *domain.OTPModel) error {
	sqlBuild := sq.Insert("otp").
		Columns("code", "otp", "id").
		Values(otp.Code, otp.OTP, otp.ID)

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("create otp: cannot build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, arg...)
	if err != nil {
		return fmt.Errorf("create otp: cannot execute query: %w", err)
	}

	return nil
}

func (r *otpRepository) Delete(ctx context.Context, code string) error {
	sqlBuild := sq.Delete("otp").Where(sq.Eq{"code": code})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("delete otp: cannot build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, arg...)
	if err != nil {
		return fmt.Errorf("delete otp: cannot execute query: %w", err)
	}

	return nil
}
