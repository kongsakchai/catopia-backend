package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type sessionRepository struct {
	db *db.Database
}

func NewSessionRepository(db *db.Database) domain.SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) Create(ctx context.Context, session *domain.Session) error {
	sqlBuild := sq.Insert("sessions").
		Columns("id", "user_id", "token", "expired_at").
		Values(sq.Expr(":id, :user_id, :token, :expired_at"))

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return errs.NewError(errs.ErrSessionCreate, err)
	}

	_, err = r.db.NamedExecContext(ctx, query, session)
	if err != nil {
		return errs.NewError(errs.ErrSessionCreate, err)
	}

	return nil
}

func (r *sessionRepository) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	var session domain.Session

	sqlBuild := sq.Select("*").
		From("sessions").
		Where(sq.Eq{"id": id})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrSessionGet, err)
	}

	err = r.db.GetContext(ctx, &session, query, arg...)
	if err != nil {
		return nil, errs.NewError(errs.ErrSessionGet, err)
	}

	return &session, nil
}

func (r *sessionRepository) Delete(ctx context.Context, id string) error {
	sqlBuild := sq.Delete("sessions").
		Where(sq.Eq{"id": id})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return errs.NewError(errs.ErrSessionDelete, err)
	}

	_, err = r.db.ExecContext(ctx, query, arg...)
	if err != nil {
		return errs.NewError(errs.ErrSessionDelete, err)
	}

	return nil
}

func (r *sessionRepository) ClearExpired(ctx context.Context) error {
	sqlBuild := sq.Delete("sessions").
		Where(" expired_at < NOW()")

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return errs.NewError(errs.ErrSessionDelete, err)
	}

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		return errs.NewError(errs.ErrSessionDelete, err)
	}

	return nil
}
