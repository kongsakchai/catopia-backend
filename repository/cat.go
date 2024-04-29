package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
)

type catRepository struct {
	db *db.Database
}

func NewCatRepository() domain.CatRepository {
	db := db.GetDB()
	return &catRepository{db}
}

func (r *catRepository) GetByID(ctx context.Context, id int, userID int) (*domain.CatModel, error) {
	sqlBuild := sq.Select("*").From("cats").Where(sq.Eq{"id": id, "user_id": userID})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get cat by id: cannot build query: %w", err)
	}

	cat := &domain.CatModel{}
	err = r.db.GetContext(ctx, cat, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get cat by id: cannot execute query: %w", err)
	}

	return cat, nil
}

func (r *catRepository) GetByUserID(ctx context.Context, userID int) ([]domain.CatModel, error) {
	sqlBuild := sq.Select("*").From("cats").Where(sq.Eq{"user_id": userID})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get cat by user id: cannot build query: %w", err)
	}

	cats := []domain.CatModel{}
	err = r.db.SelectContext(ctx, &cats, query, arg...)
	if err != nil {
		return nil, fmt.Errorf("get cat by user id:  cannot execute query: %w", err)
	}

	return cats, nil
}

func (r *catRepository) Create(ctx context.Context, cat *domain.CatModel) error {
	sqlBuild := sq.Insert("cats").
		Columns("name", "gender", "profile", "date", "user_id", "weight").
		Values(sq.Expr(":name,:gender,:profile,:date,:user_id,:weight"))

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("create cat: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.NamedExecContext(ctx, query, cat)
		if err != nil {
			return fmt.Errorf("create cat: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *catRepository) Update(ctx context.Context, cat *domain.CatModel) error {
	sqlBuild := sq.Update("cats").
		Set("name", cat.Name).
		Set("gender", cat.Gender).
		Set("profile", cat.Profile).
		Set("date", cat.Date).
		Set("weight", cat.Weight).
		Where(sq.Eq{"id": cat.ID, "user_id": cat.UserID})

	query, args, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("update cat: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("update cat: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *catRepository) Delete(ctx context.Context, id int, userID int) error {
	sqlBuild := sq.Delete("cats").Where(sq.Eq{"id": id, "user_id": userID})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("delete cat: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, arg...)
		if err != nil {
			return fmt.Errorf("delete cat: cannot execute query: %w", err)
		}

		return nil
	})
}
