package controller

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/labstack/echo/v4"

	"github.com/markable/internal/domain"
	"github.com/markable/internal/http/transport"
)

type PatientController struct {
	ps domain.PatientService
}

func NewPatientController(ps domain.PatientService) PatientController {
	return PatientController{ps: ps}
}

// FindPatientById finds a patient by ID.
//
//	@Summary		Find a patient by ID
//	@Description	Find a patient based on the provided ID
//	@Tags			Patient
//	@ID				findPatientByID
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Param			id				path		string	true	"Patient ID"
//	@Success		200				{object}	domain.BaseResponse{data=domain.Patient}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/patients/{id} [get]
func (c PatientController) FindPatientById(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	patient, err := c.ps.FindByID(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, patient)

}

// FindAllPatients retrieves all patients.
//
//	@Summary		Find all patients
//	@Description	Retrieve a list of all patients
//	@Tags			Patient
//	@ID				findAllPatients
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Success		200				{object}	domain.BaseResponse{data=[]domain.Patient}
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/patients [get]
func (c PatientController) FindAllPatients(ctx echo.Context) error {
	result, err := c.ps.FindAll()
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// CreatePatient creates a new patient.
//
//	@Summary		Create a new patient
//	@Description	Creates a patient with the provided input
//	@Tags			Patient
//	@ID				createPatient
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			body			body		domain.CreatePatientInput	true	"Patient data"
//	@Success		201				{object}	domain.BaseResponse{data=domain.Patient}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/patients [post]
func (c PatientController) CreatePatient(ctx echo.Context) error {
	var in domain.CreatePatientInput
	transport.DecodeAndValidateRequestBody(ctx, &in)
	result, err := c.ps.Create(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusCreated, result)
}

// UpdatePatient updates a patient by ID.
//
//	@Summary		Update a patient by ID
//	@Description	Updates the patient based on the provided ID and input data
//	@Tags			Patient
//	@ID				updatePatient
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			id				path		string						true	"Patient ID"
//	@Param			body			body		domain.UpdatePatientInput	true	"Updated patient data"
//	@Success		200				{object}	domain.BaseResponse{data=domain.Patient}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/patients/{id} [put]
func (c PatientController) UpdatePatient(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	var in domain.UpdatePatientInput
	transport.DecodeAndValidateRequestBody(ctx, &in)
	result, err := c.ps.Update(id, in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// DeletePatient deletes a patient by ID.
//
//	@Summary		Delete a patient by ID
//	@Description	Deletes a patient based on the provided ID
//	@Tags			Patient
//	@ID				deletePatient
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Param			id				path		string	true	"Patient ID"
//	@Success		204				{object}	domain.BaseResponse{}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/patients/{id} [delete]
func (c PatientController) DeletePatient(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = c.ps.Delete(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusNoContent, nil)
}
