package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	// Patent defines the model of a patent
	Patient struct {
		Base
		FirstName *string `json:"first_name,omitempty" db:"first_name" example:"Raj"`
		LastName  *string `json:"last_name,omitempty" db:"last_name" example:"Singhaniya"`
		Age       int64   `json:"age" db:"age" example:"26"`
		Email     *string `json:"email,omitempty" db:"email" example:"raj.singhaniya@gmail.com"`
		Phone     string  `json:"phone" db:"phone" example:"+91 9876543210"`
		Disease   string  `json:"disease" db:"disease" example:"Diabetes"`
		Address   Address `json:"address,omitempty" db:"address" sql:"jsonb"`
		BaseAudit
	} // @name Patient

)

type (
	// CreatePatientInput defines the model for CreatePatientInput
	CreatePatientInput struct {
		FirstName string  `json:"first_name,omitempty" example:"Raj"`
		LastName  string  `json:"last_name,omitempty" example:"Singhaniya"`
		Age       int64   `json:"age" example:"26"`
		Email     string  `json:"email,omitempty" example:"raj.singhaniya@gmail.com"`
		Phone     string  `json:"phone" example:"+91 9876543210"`
		Disease   string  `json:"disease" example:"Diabetes"`
		Address   Address `json:"address,omitempty"`
	} // @name CreatePatientInput

	// UpdatePatientInput defines the model for UpdatePatientInput
	UpdatePatientInput struct {
		CreatePatientInput
	} // @name UpdatePatientInput

)

type (
	// PatientRepository defines methods that any patient repository should implement
	PatientRepository interface {
		// FindByID returns a patient by id
		FindByID(ctx context.Context, id uuid.UUID) (result Patient, err error)
		// FindAll returns all patients
		FindAll(ctx context.Context) (result []Patient, err error)
		// Create creates a new patient
		Create(ctx context.Context, entity *Patient) (err error)
		// Update updates a patient
		Update(ctx context.Context, entity *Patient) (err error)
		// Delete deletes a patient
		Delete(ctx context.Context, id uuid.UUID) (err error)
	} // @name PatientRepository

	// PatientService defines the methods that any patient service should implement
	PatientService interface {
		// FindByID returns a patient by id
		FindByID(id uuid.UUID) (result Patient, err error)
		// FindAll returns all patients
		FindAll() (result []Patient, err error)
		// Create creates a new patient
		Create(in CreatePatientInput, role string) (result Patient, err error)
		// Update updates a patient
		Update(id uuid.UUID, in UpdatePatientInput, role string) (result Patient, err error)
		// Delete deletes a patient
		Delete(id uuid.UUID, role string) (err error)
	} // @name PatientService
)
