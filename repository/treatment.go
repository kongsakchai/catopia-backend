package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
)

type treatmentRepository struct {
	db *db.Database
}

func NewTreatmentRepository() domain.TreatmentRepository {
	db := db.GetDB()
	return &treatmentRepository{db}
}

func (r *treatmentRepository) GetType(ctx context.Context) ([]domain.TreatmentTypeModel, error) {
	sqlBuild := sq.Select("*").From("treatment_type")

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get treatment type: cannot build query: %w", err)
	}

	var treatmentType []domain.TreatmentTypeModel
	err = r.db.SelectContext(ctx, &treatmentType, query)
	if err != nil {
		return nil, fmt.Errorf("get treatment type: cannot execute query: %w", err)
	}

	return treatmentType, nil
}

func (r *treatmentRepository) GetByID(ctx context.Context, id int, catID int) (*domain.TreatmentModel, error) {
	sqlBuild := sq.Select("treatments.*", "treatment_type.treatment_type as name").
		From("treatments").
		LeftJoin("treatment_type ON treatments.treatment_type_id = treatment_type.id").
		Where(sq.Eq{"treatments.id": id, "cat_id": catID})

	query, arges, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get treatment by id: cannot build query: %w", err)
	}

	var treatment domain.TreatmentModel
	err = r.db.GetContext(ctx, &treatment, query, arges...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get treatment by id: cannot execute query: %w", err)
	}

	return &treatment, nil
}

func (r *treatmentRepository) GetByCatID(ctx context.Context, catID int) ([]domain.TreatmentModel, error) {
	sqlBuild := sq.Select("treatments.*", "treatment_type.treatment_type as name").
		From("treatments").
		LeftJoin("treatment_type ON treatments.treatment_type_id = treatment_type.id").
		Where(sq.Eq{"cat_id": catID})

	query, arges, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get treatment by cat id: cannot build query: %w", err)
	}

	var treatments []domain.TreatmentModel
	err = r.db.SelectContext(ctx, &treatments, query, arges...)
	if err != nil {
		return nil, fmt.Errorf("get treatment by cat id: cannot execute query: %w", err)
	}

	return treatments, nil
}

func (r *treatmentRepository) Create(ctx context.Context, treatment *domain.TreatmentModel) error {
	sqlBuild := sq.Insert("treatments").
		Columns("cat_id", "treatment_type_id", "weight", "date", "location", "vet", "detail").
		Values(treatment.CatID, treatment.TreatmentTypeID, treatment.Weight, treatment.Date, treatment.Location, treatment.Vet, treatment.Detail)

	query, arges, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("create treatment: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, arges...)
		if err != nil {
			return fmt.Errorf("create treatment: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *treatmentRepository) Update(ctx context.Context, treatment *domain.TreatmentModel) error {
	sqlBuild := sq.Update("treatments").
		Set("treatment_type_id", treatment.TreatmentTypeID).
		Set("weight", treatment.Weight).
		Set("date", treatment.Date).
		Set("location", treatment.Location).
		Set("vet", treatment.Vet).
		Set("detail", treatment.Detail).
		Where(sq.Eq{"id": treatment.ID, "cat_id": treatment.CatID})

	query, arges, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("update treatment: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, arges...)
		if err != nil {
			return fmt.Errorf("update treatment: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *treatmentRepository) Delete(ctx context.Context, id int, catID int) error {
	fmt.Println(id, catID)

	sqlBuild := sq.Delete("treatments").Where(sq.Eq{"id": id, "cat_id": catID})

	query, arges, err := sqlBuild.ToSql()
	if err != nil {
		return err
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err = tx.ExecContext(ctx, query, arges...)
		if err != nil {
			return err
		}

		return nil
	})
}
