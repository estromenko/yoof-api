package services

import (
	"github.com/estromenko/yoof-api/internal/models"
	"github.com/estromenko/yoof-api/internal/repo"
)

var (
	trains = map[int]float32{
		0: 1.2,

		1: 1.375,
		2: 1.375,
		3: 1.375,

		4: 1.55,
		5: 1.55,

		6: 1.725,
		7: 1.725,

		-1: 1.9, // Постоянная физ. нагрузка и тп
	}
)

type CalcService struct {
	repo *repo.DishRepo
}

func NewCalcService(repo *repo.DishRepo) *CalcService {
	return &CalcService{
		repo: repo,
	}
}

func (s *CalcService) Repo() *repo.DishRepo {
	return s.repo
}

func (s *CalcService) findTrainingCoefficient(number int) float32 {
	return trains[number]
}

func (s *CalcService) CountCalories(gender int, weight float32, growth float32, training int) float32 {
	calories := 10*weight + 6.25*growth
	if gender == 0 {
		calories = calories - 161
	} else {
		calories = calories + 5
	}
	calories = calories * s.findTrainingCoefficient(training)
	return calories
}

func (s *CalcService) GetDayList(eatingAmount int, calories float32) []float32 {
	var dayList []float32

	if eatingAmount == 3 {
		dayList = []float32{30.0, 0.0, 40.0, 0.0, 30.0}
	} else if eatingAmount == 5 {
		dayList = []float32{20.0, 10.0, 35.0, 10.0, 25.0}
	} else {
		return []float32{}
	}
	for i := range dayList {
		dayList[i] = dayList[i] * calories / 100.0
	}
	return dayList
}

func (s *CalcService) getDiffer(calories float32, differ float32) (float32, float32) {
	lte := calories * (1.0 + differ/100.0)
	gte := calories * (1.0 - differ/100.0)
	return lte, gte
}

func (s *CalcService) GetRandomDishByCalories(calories float32, day_time int, differ float32) *models.Dish {
	lte, gte := s.getDiffer(calories, differ)
	dishes, _ := s.repo.FindAllInCaloriesRange(gte, lte)
	dish := dishes[0]
	return dish
}
