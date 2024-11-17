//go:build wireinject

package dependency

import (
	"github.com/google/wire"
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

func NewConfig(opt config.Options) (config.MarkAbleConfig, error) {
	wire.Build(
		config.NewConfig, // This should call your NewConfig function which takes Options
	)
	return config.MarkAbleConfig{}, nil
}

func NewDatabaseConfig(cfg config.MarkAbleConfig) (*pgxpool.Pool, error) {
	wire.Build(
		database.NewDB,
	)
	return &pgxpool.Pool{}, nil
}

func NewMarkableApi(cfg config.MarkAbleConfig, db *pgxpool.Pool) (api.MarkAbleApi, error) {
	wire.Build(
		util.NewAppUtil,
		security.NewJwtSecurityManager,
		repository.NewTransactioner,
		repository.NewUserRepository,
		repository.NewPatientRepository,
		service.NewUserService,
		service.NewPatientService,
		controller.NewUserController,
		controller.NewPatientController,
		api.NewMarkableApi,
	)
	return api.MarkAbleApi{}, nil
}
