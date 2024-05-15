package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type treatmentRepository struct {
	db *db.Database
}

func NewTreatmentRepository(db *db.Database) domain.TreatmentRepository {
	return &treatmentRepository{db}
}
func (r *treatmentRepository) GetType(ctx context.Context) ([]domain.TreatmentType, error) {
	getSql, _, err := sq.Select("*").From("treatment_type").ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentTypeGet, err)
	}

	var treatmentType []domain.TreatmentType
	err = r.db.SelectContext(ctx, &treatmentType, getSql)
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentTypeGet, db.HandlerError(err))
	}

	return treatmentType, nil
}

func (r *treatmentRepository) GetByID(ctx context.Context, id int64, catID int64) (*domain.Treatment, error) {
	getSql, args, err := sq.Select("t.*", "tt.treatment_type as name").From("treatment t").
		LeftJoin("treatment_type tt ON t.treatment_type_id = tt.id").
		Where(sq.Eq{"t.id": id, "t.cat_id": catID}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetByID, err)
	}

	treatment := &domain.Treatment{}
	err = r.db.GetContext(ctx, treatment, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetByID, db.HandlerError(err))
	}

	return treatment, nil
}

func (r *treatmentRepository) GetByCatID(ctx context.Context, catID int64) ([]domain.Treatment, error) {
	getSql, args, err := sq.Select("t.*", "tt.treatment_type as name").From("treatment t").
		LeftJoin("treatment_type tt ON t.treatment_type_id = tt.id").
		Where(sq.Eq{"cat_id": catID}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetByCatID, err)
	}

	var treatment []domain.Treatment
	err = r.db.SelectContext(ctx, &treatment, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetByCatID, db.HandlerError(err))
	}

	return treatment, nil
}

func (r *treatmentRepository) Create(ctx context.Context, treatment *domain.Treatment) error {
	insertSql, args, err := sq.Insert("treatment").
		Columns("cat_id", "treatment_type_id", "date", "location", "vet", "detail", "appointment", "appointment_date").
		Values(treatment.CatID, treatment.TreatmentTypeID, treatment.Date, treatment.Location, treatment.Vet, treatment.Detail, treatment.Appointment, treatment.AppointmentDate).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrTreatmentCreate, err)
	}

	_, err = r.db.ExecContext(ctx, insertSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrTreatmentCreate, db.HandlerError(err))
	}

	return nil
}

func (r *treatmentRepository) Update(ctx context.Context, treatment *domain.Treatment) error {
	updateSql, args, err := sq.Update("treatment").
		Set("treatment_type_id", treatment.TreatmentTypeID).
		Set("date", treatment.Date).
		Set("location", treatment.Location).
		Set("vet", treatment.Vet).
		Set("detail", treatment.Detail).
		Set("appointment", treatment.Appointment).
		Set("appointment_date", treatment.AppointmentDate).
		Where(sq.Eq{"id": treatment.ID, "cat_id": treatment.CatID}).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrTreatmentUpdate, err)
	}

	_, err = r.db.ExecContext(ctx, updateSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrTreatmentUpdate, db.HandlerError(err))
	}

	return nil
}

func (r *treatmentRepository) Delete(ctx context.Context, id int64) error {
	deleteSql, args, err := sq.Delete("treatment").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return errs.NewError(errs.ErrTreatmentDelete, err)
	}

	_, err = r.db.ExecContext(ctx, deleteSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrTreatmentDelete, db.HandlerError(err))
	}

	return nil
}

func (r *treatmentRepository) GetTreatmentNoti(ctx context.Context, userID int64) ([]domain.TreatmentNoti, error) {
	getSql, args, err := sq.Select("t.appointment_date", "t.appointment", "c.name as name", "c.id as cat_id", "t.id as id").From("treatment t").
		Join("cat c ON t.cat_id = c.id").
		Where(sq.And{sq.Eq{"c.user_id": userID}, sq.Expr("t.appointment_date BETWEEN CURDATE() AND DATE_ADD(CURDATE(),INTERVAL 7 DAY)")}).
		ToSql()

	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetNoti, err)
	}

	var treatment []domain.TreatmentNoti
	err = r.db.SelectContext(ctx, &treatment, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrTreatmentGetNoti, db.HandlerError(err))
	}

	return treatment, nil
}
