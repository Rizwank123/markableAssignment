// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markable/internal/database"
	"github.com/markable/internal/http/api"
	"github.com/markable/internal/http/controller"
	"github.com/markable/internal/pkg/config"
	"github.com/markable/internal/pkg/security"
	"github.com/markable/internal/pkg/util"
	"github.com/markable/internal/repository"
	"github.com/markable/internal/service"
)

// Injectors from wire.go:

func NewConfig(opt config.Options) (config.MarkAbleConfig, error) {
	markAbleConfig, err := config.NewConfig(opt)
	if err != nil {
		return config.MarkAbleConfig{}, err
	}
	return markAbleConfig, nil
}

func NewDatabaseConfig(cfg config.MarkAbleConfig) (*pgxpool.Pool, error) {
	pool := database.NewDB(cfg)
	return pool, nil
}

func NewMarkableApi(cfg config.MarkAbleConfig, db *pgxpool.Pool) (*api.MarkAbleApi, error) {
	patientRepository := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)
	appUtil := util.NewAppUtil()
	manager := security.NewJwtSecurityManager(cfg)
	transactioner := repository.NewTransactioner(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(appUtil, cfg, manager, transactioner, userRepository)
	userController := controller.NewUserController(userService)
	markAbleApi := api.NewMarkableApi(cfg, patientController, userController)
	return markAbleApi, nil
}
