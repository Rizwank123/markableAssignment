package integration

import (
	"net/http"
	"testing"

	"github.com/markable/internal/domain"
	"github.com/markable/tests/helper"
)

func TestCreatePatient(t *testing.T) {
	t.Run("Create New Patient", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		// create and request login
		loginInput := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		lrec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, loginInput)
		if err != nil {
			t.Errorf("Login Failed : %v", err)
		}
		// Parse the response
		var loginResponse domain.BaseResponse
		helper.ParseResponse(t, lrec, &loginResponse)
		// Parse the Entity
		var loginOutput domain.LoginOutput
		helper.ParseEntityData(t, loginResponse.Data, &loginOutput)
		token := make(map[string]string)
		token["Authorization"] = "Bearer " + loginOutput.Token
		// Create a new patient
		patient := domain.CreatePatientInput{
			FirstName: "Aman",
			LastName:  "Singh",
			Email:     "aman.singh@example.com",
			Phone:     "+91 9876543210",
			Disease:   "Diabetes",
			Age:       24,
			Address: domain.Address{
				Street:  "123, ABC Street",
				City:    "New York",
				State:   "New York",
				Country: "USA",
				Pincode: "10001",
			},
		}
		patientRce, err := helper.SendRequest(e, tApi.PatientController.CreatePatient, http.MethodPost, "/patients", nil, token, patient)
		if err != nil {
			t.Errorf("Send Request Failed : %v", err)
		}
		// Parse the response
		var patientResponse domain.BaseResponse
		helper.ParseResponse(t, patientRce, &patientResponse)
		// Check the response
		codeWant := http.StatusCreated
		codeGot := patientRce.Code
		if codeWant != codeGot {
			t.Errorf("Expected %d, got %d", codeWant, codeGot)
		}
		// Parse the Entity
		var patientOutput domain.Patient
		helper.ParseEntityData(t, patientResponse.Data, &patientOutput)
		// Check the response
		if patientOutput.ID.IsNil() {
			t.Errorf("Patient ID is nil")
		}
		firstNmeWant := "Aman"
		lastNameWant := "Singh"
		firstNameGot := *patientOutput.FirstName
		lastNameGot := *patientOutput.LastName
		if firstNameGot != firstNmeWant {
			t.Errorf("Expected %s, got %s", firstNmeWant, firstNameGot)
		}
		if lastNameWant != lastNameGot {
			t.Errorf("Expected %s, got %s", lastNameWant, lastNameGot)
		}
		emailWant := "aman.singh@example.com"
		emailGot := *patientOutput.Email
		if emailWant != emailGot {
			t.Errorf("Expected %s, got %s", emailWant, emailGot)
		}
		phoneWant := "+91 9876543210"
		if patientOutput.Phone != phoneWant {
			t.Errorf("Expected %s, got %s", phoneWant, patientOutput.Phone)
		}

	})

}
func TestFindAllPatients(t *testing.T) {
	t.Run("Find all Patients", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		// create and request login
		loginInput := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		lrec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, loginInput)
		if err != nil {
			t.Errorf("Login Failed : %v", err)
		}
		// Parse the response
		var loginResponse domain.BaseResponse
		helper.ParseResponse(t, lrec, &loginResponse)
		// Parse the Entity
		var loginOutput domain.LoginOutput
		helper.ParseEntityData(t, loginResponse.Data, &loginOutput)
		token := make(map[string]string)
		token["Authorization"] = "Bearer " + loginOutput.Token
		rec, err := helper.SendRequest(e, tApi.PatientController.FindAllPatients, http.MethodGet, "/patients", nil, token, nil)
		if err != nil {
			t.Errorf("Failed to find all patients : %v", err)
		}
		// Parse the response
		var patientsResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &patientsResponse)
		// Parse the Entity
		var patientsOutput []domain.Patient
		helper.ParseEntityData(t, patientsResponse.Data, &patientsOutput)
		if len(patientsOutput) < 1 {
			t.Errorf("Expected at least one patient, got %d", len(patientsOutput))
		}

	})
}

func TestFindPatientById(t *testing.T) {
	t.Run("Find Patient by ID", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		// create and request login
		loginInput := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		lrec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, loginInput)
		if err != nil {
			t.Errorf("Login Failed : %v", err)
		}
		// Parse the response
		var loginResponse domain.BaseResponse
		helper.ParseResponse(t, lrec, &loginResponse)
		// Parse the Entity
		var loginOutput domain.LoginOutput
		helper.ParseEntityData(t, loginResponse.Data, &loginOutput)
		token := make(map[string]string)
		token["Authorization"] = "Bearer " + loginOutput.Token
		rec, err := helper.SendRequest(e, tApi.PatientController.FindAllPatients, http.MethodGet, "/patients", nil, token, nil)
		if err != nil {
			t.Errorf("Failed to find all patients : %v", err)
		}
		// Parse the response
		var patientsResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &patientsResponse)
		// Parse the Entity
		var patientsOutput []domain.Patient
		helper.ParseEntityData(t, patientsResponse.Data, &patientsOutput)
		if len(patientsOutput) < 1 {
			t.Errorf("Expected at least one patient, got %d", len(patientsOutput))
		}
		pathParams := map[string]string{}
		pathParams["id"] = patientsOutput[0].ID.String()
		// Find Patient by ID
		rec, err = helper.SendRequest(e, tApi.PatientController.FindPatientById, http.MethodPost, "/patients/"+patientsOutput[0].ID.String(), pathParams, token, nil)
		if err != nil {
			t.Errorf("Failed to find patient by ID : %v", err)
		}
		// Parse the response
		var patientResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &patientResponse)
		// Parse the Entity
		var patientOutput domain.Patient
		helper.ParseEntityData(t, patientResponse.Data, &patientOutput)
		if patientOutput.ID.String() != pathParams["id"] {
			t.Errorf("Expected patient ID %s, got %s", pathParams["id"], patientOutput.ID.String())
		}
	})
}

func TestUpdatePatient(t *testing.T) {
	// Create a new test environment
	t.Run("Update Patient by ID", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		// create and request login
		loginInput := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		lrec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, loginInput)
		if err != nil {
			t.Errorf("Login Failed : %v", err)
		}
		// Parse the response
		var loginResponse domain.BaseResponse
		helper.ParseResponse(t, lrec, &loginResponse)
		// Parse the Entity
		var loginOutput domain.LoginOutput
		helper.ParseEntityData(t, loginResponse.Data, &loginOutput)
		token := make(map[string]string)
		token["Authorization"] = "Bearer " + loginOutput.Token
		rec, err := helper.SendRequest(e, tApi.PatientController.FindAllPatients, http.MethodGet, "/patients", nil, token, nil)
		if err != nil {
			t.Errorf("Failed to find all patients : %v", err)
		}
		// Parse the response
		var patientsResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &patientsResponse)
		// Parse the Entity
		var patientsOutput []domain.Patient
		helper.ParseEntityData(t, patientsResponse.Data, &patientsOutput)
		if len(patientsOutput) < 1 {
			t.Errorf("Expected at least one patient, got %d", len(patientsOutput))
		}
		pathParams := map[string]string{}
		pathParams["id"] = patientsOutput[0].ID.String()
		// Update Patient
		updateInput := domain.UpdatePatientInput{
			CreatePatientInput: domain.CreatePatientInput{
				FirstName: "James",
				LastName:  "Bond",
				Email:     "james.bond@example.com",
				Age:       30,
			},
		}
		rec, err = helper.SendRequest(e, tApi.PatientController.UpdatePatient, http.MethodPut, "/patients/"+patientsOutput[0].ID.String(), pathParams, nil, updateInput)
		if err != nil {
			t.Errorf("Failed to update patient : %v", err)
		}
		// Parse the response
		var updateResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &updateResponse)
		// Parse the Entity
		var updateOutput domain.Patient
		helper.ParseEntityData(t, updateResponse.Data, &updateOutput)
		firstNameWant := updateInput.FirstName
		firstNameGot := *updateOutput.FirstName
		if firstNameWant != firstNameGot {
			t.Errorf("Expected firstName to be %s, got %s", firstNameWant, firstNameGot)
		}
		lastNameWant := updateInput.LastName
		lastNameGot := *updateOutput.LastName
		if lastNameWant != lastNameGot {
			t.Errorf("Expected lastName to be %s, got %s", lastNameWant, lastNameGot)
		}
		emailWant := updateInput.Email
		emailGot := *updateOutput.Email
		if emailWant != emailGot {
			t.Errorf("Expected email to be %s, got %s", emailWant, emailGot)
		}
		ageWant := updateInput.Age
		ageGot := updateOutput.Age
		if ageWant != ageGot {
			t.Errorf("Expected age to be %d, got %d", ageWant, ageGot)
		}
	})

}

func TestDeletePatient(t *testing.T) {
	t.Run("Delete Patient", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		// create and request login
		loginInput := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		lrec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, loginInput)
		if err != nil {
			t.Errorf("Login Failed : %v", err)
		}
		// Parse the response
		var loginResponse domain.BaseResponse
		helper.ParseResponse(t, lrec, &loginResponse)
		// Parse the Entity
		var loginOutput domain.LoginOutput
		helper.ParseEntityData(t, loginResponse.Data, &loginOutput)
		token := make(map[string]string)
		token["Authorization"] = "Bearer " + loginOutput.Token
		rec, err := helper.SendRequest(e, tApi.PatientController.FindAllPatients, http.MethodGet, "/patients", nil, token, nil)
		if err != nil {
			t.Errorf("Failed to find all patients : %v", err)
		}
		// Parse the response
		var patientsResponse domain.BaseResponse
		helper.ParseResponse(t, rec, &patientsResponse)
		// Parse the Entity
		var patientsOutput []domain.Patient
		helper.ParseEntityData(t, patientsResponse.Data, &patientsOutput)
		if len(patientsOutput) < 1 {
			t.Errorf("Expected at least one patient, got %d", len(patientsOutput))
		}
		pathParams := map[string]string{}
		pathParams["id"] = patientsOutput[0].ID.String()
		rec, err = helper.SendRequest(e, tApi.PatientController.DeletePatient, http.MethodDelete, "/patients/"+patientsOutput[0].ID.String(), pathParams, token, nil)
		if err != nil {
			t.Errorf("Failed to delete patient : %v", err)
		}
		// Parse the response
		//var deleteResponse domain.BaseResponse
		//helper.ParseResponse(t, rec, &deleteResponse)
		// Check the status code
		codeWant := http.StatusNoContent
		codeGot := rec.Code
		if codeWant != codeGot {
			t.Errorf("Expected status code %d, got %d", codeWant, codeGot)
		}
	})
}
