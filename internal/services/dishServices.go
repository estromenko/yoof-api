package services

import (
	"fmt"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/estromenko/yoof-api/internal/repo"
)

type DishServiceConfig struct{}

type DishService struct {
	Config *DishServiceConfig
	repo   *repo.DishRepo
}

func NewDishService(repo *repo.DishRepo, config *DishServiceConfig) *DishService {
	return &DishService{
		repo:   repo,
		Config: config,
	}
}

func (s *DishService) Repo() *repo.DishRepo {
	return s.repo
}

func (s *DishService) validate(dish *models.Dish) string {
	// TODO
	return ""
}

func (s *DishService) Create(dish *models.Dish) error {
	errors := s.validate(dish)
	if errors != "" {
		return fmt.Errorf(errors)
	}

	return s.repo.Create(dish)
}
