package services

import (
	"github.com/estromenko/yoof-api/internal/repo"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	User *UserServiceConfig `mapstructure:"user"`
}

type Services struct {
	UserService *UserService
}

func InitServices(conn *sqlx.DB, config *Config) *Services {
	return &Services{
		UserService: NewUserService(
			repo.NewUserRepo(conn),
			config.User,
		),
	}
}
