package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	SourceKey = "CONFIG_SOURCE"
	SourceEnv = "ENVIRONMENT"
)

type MarkAbleConfig struct {
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_DATABASE_NAME"`

	AppPort          int    `mapstructure:"APP_PORT"`
	AuthSecret       string `mapstructure:"AUTH_SECRET"`
	AuthExpiryPeriod int    `mapstructure:"AUTH_EXPIRY_PERIOD"`

	SwaggerHostUrl    string `mapstructure:"SWAGGER_HOST_URL"`
	SwaggerHostScheme string `mapstructure:"SWAGGER_HOST_SCHEME"`
	SwaggerUsername   string `mapstructure:"SWAGGER_USERNAME"`
	SwaggerPassword   string `mapstructure:"SWAGGER_PASSWORD"`
}

type Options struct {
	ConfigFile       string
	ConfigFileSource string
}

// NewConfig creates a new MarkAbleConfig by reading the provided options.
func NewConfig(opt Options) (MarkAbleConfig, error) {
	return NewFromEnvironmentVariable(opt)
}

// NewFromEnvironmentVariable loads configuration from the environment variables.
func NewFromEnvironmentVariable(opt Options) (MarkAbleConfig, error) {
	viper.SetConfigFile(opt.ConfigFile)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return MarkAbleConfig{}, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg MarkAbleConfig
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return MarkAbleConfig{}, fmt.Errorf("failed to load configuration: %v", err)
	}

	return cfg, nil
}