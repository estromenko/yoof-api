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
	CalcService *CalcService
}

func InitServices(conn *sqlx.DB, config *Config) *Services {
	userRepo := repo.NewUserRepo(conn)
	dishRepo := repo.NewDishRepo(conn)
	return &Services{
		UserService: NewUserService(
			userRepo,
			config.User,
		),
		DishService: NewDishService(
			dishRepo,
			config.Dish,
		),
		CalcService: NewCalcService(
			dishRepo,
		),
	}
}
