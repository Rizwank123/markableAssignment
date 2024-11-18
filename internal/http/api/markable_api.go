package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/markable/internal/http/controller"
	"github.com/markable/internal/pkg/config"
)

type MarkAbleApi struct {
	cfg               config.MarkAbleConfig
	PatientController controller.PatientController
	UserController    controller.UserController
}

// NewMarkableApi creates a new MarkableApi instance
//
//	@title						Markable API
//	@version					1.0
//	@description				Markable application's set of APIs
//	@termsOfService				https://example.com/terms
//	@contact.name				Mohammad Developer
//	@contact.url				https://rizwank123.github.io
//	@contact.email				md.rizwank431@gmail.com
//	@host						localhost:7700
//	@BasePath					/api/v1
//	@schemes					http https
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
func NewMarkableApi(cfg config.MarkAbleConfig, pc controller.PatientController, uc controller.UserController) *MarkAbleApi {
	return &MarkAbleApi{
		cfg:               cfg,
		PatientController: pc,
		UserController:    uc,
	}
}

func (b MarkAbleApi) SetupRoutes(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	auth := echojwt.JWT([]byte(b.cfg.AuthSecret))

	userApi := apiV1.Group("/users")
	userApi.POST("/login", b.UserController.Login)
	userApi.POST("", b.UserController.RegisterUser)
	secureApi := apiV1.Group("/users")
	secureApi.Use(auth)
	secureApi.GET("/:id", b.UserController.FindByID)

	patientApi := apiV1.Group("/patients")
	patientApi.GET("", b.PatientController.FindAllPatients)
	patientApi.GET("/:id", b.PatientController.FindPatientById)
	patientApi.POST("", b.PatientController.CreatePatient)
	patientApi.PUT("/:id", b.PatientController.UpdatePatient)
	patientApi.DELETE("/:id", b.PatientController.DeletePatient)

}
