package services

import (
	"github.com/estromenko/yoof-api/internal/repo"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	User *UserServiceConfig `mapstructure:"user"`
	Dish *DishServiceConfig `mapstructure:"dish"`
}

type Services struct {
	UserService *UserService
	DishService *DishService
}

func InitServices(conn *sqlx.DB, config *Config) *Services {
	return &Services{
		UserService: NewUserService(
			repo.NewUserRepo(conn),
			config.User,
		),
		DishService: NewDishService(
			repo.NewDishRepo(conn),
			config.Dish,
		),
	}
}
