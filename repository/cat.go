package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type catRepository struct {
	db *db.Database
}

func NewCatRepository(db *db.Database) domain.CatRepository {
	return &catRepository{db}
}

func (r *catRepository) GetByID(ctx context.Context, id int64, userID int64) (*domain.Cat, error) {
	getSql, args, err := squirrel.Select("*").From("cat").Where(squirrel.Eq{"id": id, "user_id": userID}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, err)
	}

	var cat domain.Cat
	err = r.db.GetContext(ctx, &cat, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return &cat, nil
}

func (r *catRepository) GetByUserID(ctx context.Context, userID int64) ([]domain.Cat, error) {
	getSql, args, err := squirrel.Select("*").From("cat").Where(squirrel.Eq{"user_id": userID}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, err)
	}

	var cats []domain.Cat
	err = r.db.SelectContext(ctx, &cats, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return cats, nil
}

func (r *catRepository) GetByUserIDs(ctx context.Context, userID []int64) ([]domain.Cat, error) {
	getSql, args, err := squirrel.Select("*").From("cat").Where(squirrel.Eq{"user_id": userID}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, err)
	}

	var cats []domain.Cat
	err = r.db.SelectContext(ctx, &cats, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return cats, nil
}

func (r *catRepository) Create(ctx context.Context, cat *domain.Cat) error {
	createSql, _, err := squirrel.Insert("cat").
		Columns("name", "weight", "gender", "profile", "date", "breeding", "aggression", "shyness", "extraversion", "user_id").
		Values(squirrel.Expr(":name, :weight, :gender, :profile, :date, :breeding, :aggression, :shyness, :extraversion, :user_id")).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}

	_, err = r.db.NamedExecContext(ctx, createSql, cat)
	if err != nil {
		return errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return nil
}

func (r *catRepository) Update(ctx context.Context, cat *domain.Cat) error {
	updateSql, args, err := squirrel.Update("cat").
		Set("name", cat.Name).
		Set("weight", cat.Weight).
		Set("gender", cat.Gender).
		Set("profile", cat.Profile).
		Set("date", cat.Date).
		Set("breeding", cat.Breeding).
		Set("aggression", cat.Aggression).
		Set("shyness", cat.Shyness).
		Set("extraversion", cat.Extraversion).
		Where(squirrel.Eq{"id": cat.ID, "user_id": cat.UserID}).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}

	_, err = r.db.ExecContext(ctx, updateSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return nil
}

func (r *catRepository) UpdateGroup(ctx context.Context, id int64, group int64) error {
	updateSql, args, err := squirrel.Update("cat").
		Set("group_id", group).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrCatUpdateGroup, err)
	}

	_, err = r.db.ExecContext(ctx, updateSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrCatUpdateGroup, db.HandlerError(err))
	}

	return nil
}

func (r *catRepository) Delete(ctx context.Context, id int64) error {
	deleteSql, args, err := squirrel.Delete("cat").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return errs.NewError(errs.ErrInternal, err)
	}

	_, err = r.db.ExecContext(ctx, deleteSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrInternal, db.HandlerError(err))
	}

	return nil
}

func (r *catRepository) GetBreedingByGroup(ctx context.Context, group []int64, random bool) ([]string, error) {
	build := squirrel.Select("breed").From("breeding").
		LeftJoin("cat_group ON breeding.group_name = cat_group.group_name").
		Where(squirrel.Eq{"cat_group.group": group}).Limit(10)

	if random {
		build = build.OrderBy("RAND()")
	} else {
		build = build.OrderBy("cat_group.count DESC")
	}

	findSql, args, err := build.ToSql()

	if err != nil {
		return nil, errs.NewError(errs.ErrCatGetCluster, err)
	}

	var breeds []string
	err = r.db.SelectContext(ctx, &breeds, findSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrCatGetCluster, db.HandlerError(err))
	}

	return breeds, nil
}

func (r *catRepository) GetBreedingByRandom(ctx context.Context) ([]string, error) {
	findSql, args, err := squirrel.Select("breed").From("breeding").OrderBy("RAND()").Limit(10).ToSql()

	if err != nil {
		return nil, errs.NewError(errs.ErrCatGetRandom, err)
	}

	var breeds []string
	err = r.db.SelectContext(ctx, &breeds, findSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrCatGetRandom, db.HandlerError(err))
	}

	return breeds, nil
}
