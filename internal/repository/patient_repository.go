package repository

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/markable/internal/domain"
)

type pgxPatientRepository struct {
	db *pgxpool.Pool
}

func NewPatientRepository(db *pgxpool.Pool) domain.PatientRepository {
	return &pgxPatientRepository{db: db}

}

// Create implements domain.PatientRepository.
func (r *pgxPatientRepository) Create(ctx context.Context, entity *domain.Patient) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `INSERT INTO patient(first_name, last_name, age, email, phone, disease, address) values($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	args := []interface{}{entity.FirstName, entity.LastName, entity.Age, entity.Email, entity.Phone, entity.Disease, entity.Address}

	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt)
	}
	return err

}

// Delete implements domain.PatientRepository.
func (r *pgxPatientRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `DELETE FROM patient WHERE id = $1`
	args := []interface{}{id}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}

// FindAll implements domain.PatientRepository.
func (r *pgxPatientRepository) FindAll(ctx context.Context) (result []domain.Patient, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM patient`
	var rows pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		rows, err = tx.Query(ctx, q)
	} else {
		rows, err = r.db.Query(ctx, q)
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.Patient])
	return result, err
}

// FindByID implements domain.PatientRepository.
func (r *pgxPatientRepository) FindByID(ctx context.Context, id uuid.UUID) (result domain.Patient, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM patient WHERE id = $1`
	args := []interface{}{id}
	var row pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		row, err = tx.Query(ctx, q, args...)
	} else {
		row, err = r.db.Query(ctx, q, args...)
	}
	defer row.Close()
	if err != nil {
		return result, err
	}
	result, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByNameLax[domain.Patient])
	if err != nil {
		return result, err
	}
	return result, nil

}

// Update implements domain.PatientRepository.
func (r *pgxPatientRepository) Update(ctx context.Context, entity *domain.Patient) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE patient SET first_name = $1, last_name = $2, age = $3, email = $4, phone = $5, address = $6, disease = $7 WHERE id = $8 RETURNING updated_at`
	args := []interface{}{entity.FirstName, entity.LastName, entity.Age, entity.Email, entity.Phone, entity.Address, entity.Disease, entity.ID}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&entity.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&entity.UpdatedAt)
	}
	return err

}
