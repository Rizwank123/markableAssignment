package integration

import (
	"net/http"
	"testing"

	"github.com/markable/internal/domain"
	"github.com/markable/tests/helper"
)

func TestRegister(t *testing.T) {
	t.Run("Register a new user", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		//  Create and send a request
		user := domain.RegisterUserInput{
			FullName: "John Doe",
			UserName: "john.doe@example.com",
			Password: "password123",
			Role:     "DOCTOR",
		}
		rec, err := helper.SendRequest(e, tApi.UserController.RegisterUser, http.MethodPost, "/users", nil, nil, user)
		if err != nil {
			t.Errorf("error sending request: %v", err)
		}
		// Check method and code
		if rec.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
		}
		// Parse response
		var resp domain.BaseResponse
		helper.ParseResponse(t, rec, &resp)

		// Parse and verify the data
		var entityData domain.User
		helper.ParseEntityData(t, resp.Data, &entityData)
		if entityData.ID.IsNil() {
			t.Errorf("expected ID to be set, got nil")
		}
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login a user", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		//  Create and send a request
		in := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "password123",
		}
		rec, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, in)
		if err != nil {
			t.Errorf("error sending request: %v", err)
		}
		// Check method and code
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
		}
		// Parse response
		var resp domain.BaseResponse
		helper.ParseResponse(t, rec, &resp)
		// Parse and verify the data
		var entityData domain.LoginOutput
		helper.ParseEntityData(t, resp.Data, &entityData)
		expectedExpire := 3600
		if entityData.ExpiresIn != int64(expectedExpire) {
			t.Errorf("expected ExpiresIn %d, got %d", expectedExpire, entityData.ExpiresIn)
		}
		if entityData.Token == "" {
			t.Errorf("expected Token to be set, got empty string")
		}
	})
	t.Run("Failed login", func(t *testing.T) {
		tApi, e, tearDownSuite := helper.SetupSuite(t)
		defer tearDownSuite(t)
		//  Create and send a request
		in := domain.LoginInput{
			UserName: "john.doe@example.com",
			Password: "wrongpassword",
		}
		_, err := helper.SendRequest(e, tApi.UserController.Login, http.MethodPost, "/users/login", nil, nil, in)
		// Check error message
		messageWant := "wrong password"
		if err != nil {
			if err.Error() != messageWant {
				t.Errorf("expected error message %q, got nil", messageWant)
			}
		}
	})

}
