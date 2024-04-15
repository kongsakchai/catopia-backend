package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
)

type sessionRepository struct {
	db *db.Database
}

func NewSessionRepository() domain.SessionRepository {
	db := db.GetDB()
	return &sessionRepository{db}
}

func (r *sessionRepository) Create(ctx context.Context, session *domain.Session) error {
	sqlBuild := sq.Insert("sessions").
		Columns("id", "user_id", "token", "expired_at").
		Values(sq.Expr(":id, :user_id, :token, :expired_at"))

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("create session: error building query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.NamedExecContext(ctx, query, session)
		if err != nil {
			return fmt.Errorf("create session: error executing query: %w", err)
		}

		return nil
	})
}

func (r *sessionRepository) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	var session domain.Session

	sqlBuild := sq.Select("*").
		From("sessions").
		Where(sq.Eq{"id": id})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get session by id: error building query: %w", err)
	}

	err = r.db.GetContext(ctx, &session, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get session by id: error executing query: %w", err)
	}

	return &session, nil
}

func (r *sessionRepository) Delete(ctx context.Context, id string) error {
	sqlBuild := sq.Delete("sessions").
		Where(sq.Eq{"id": id})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("delete session: error building query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, arg...)
		if err != nil {
			return fmt.Errorf("delete session: error executing query: %w", err)
		}

		return nil
	})
}
