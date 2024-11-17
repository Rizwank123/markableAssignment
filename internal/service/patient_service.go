package service

import (
	"context"
	"errors"
	"reflect"

	"github.com/gofrs/uuid/v5"

	"github.com/markable/internal/domain"
)

type patientServiceImpl struct {
	pr domain.PatientRepository
}

func NewPatientService(pr domain.PatientRepository) domain.PatientService {
	return &patientServiceImpl{pr: pr}
}

// Create implements domain.PatientService.
func (s *patientServiceImpl) Create(in domain.CreatePatientInput) (result domain.Patient, err error) {
	result = domain.Patient{
		FirstName: &in.FirstName,
		LastName:  &in.LastName,
		Email:     &in.Email,
		Phone:     in.Phone,
		Disease:   in.Disease,
		Age:       in.Age,
		Address:   in.Address,
	}
	err = s.pr.Create(context.Background(), &result)
	if err != nil {
		return result, err
	}
	return result, nil

}

// Delete implements domain.PatientService.
func (s patientServiceImpl) Delete(id uuid.UUID) (err error) {
	result, err := s.pr.FindByID(context.Background(), id)
	if err != nil {
		return err
	}
	if result.ID.IsNil() {
		return errors.New("patient not found")
	}
	err = s.pr.Delete(context.Background(), uuid.UUID(result.ID))
	return err
}

// FindAll implements domain.PatientService.
func (s *patientServiceImpl) FindAll() (result []domain.Patient, err error) {
	return s.pr.FindAll(context.Background())
}

// FindByID implements domain.PatientService.
func (s *patientServiceImpl) FindByID(id uuid.UUID) (result domain.Patient, err error) {
	return s.pr.FindByID(context.Background(), id)
}

// Update implements domain.PatientService.
func (s *patientServiceImpl) Update(id uuid.UUID, in domain.UpdatePatientInput) (result domain.Patient, err error) {
	result, err = s.pr.FindByID(context.Background(), id)
	if err != nil {
		return result, err
	}
	if result.ID.IsNil() {
		return result, errors.New("patient not found")
	}
	if in.FirstName != "" {
		result.FirstName = &in.FirstName
	}
	if in.LastName != "" {
		result.LastName = &in.LastName
	}
	if in.Age != 0 {
		result.Age = in.Age
	}
	if in.Email != "" {
		result.Email = &in.Email
	}
	if in.Phone != "" {
		result.Phone = in.Phone
	}
	if !reflect.DeepEqual(in.Address, domain.Address{}) {
		result.Address = in.Address
	}
	if in.Disease != "" {
		result.Disease = in.Disease
	}
	err = s.pr.Update(context.Background(), &result)
	if err != nil {
		return result, err
	}
	return result, err

}
