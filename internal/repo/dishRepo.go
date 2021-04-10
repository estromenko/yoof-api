package repo

import (
	"database/sql"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishRepo struct {
	db *sqlx.DB
}

func NewDishRepo(db *sqlx.DB) *DishRepo {
	return &DishRepo{
		db: db,
	}
}

func (r *DishRepo) Create(dish *models.Dish) error {
	return r.db.QueryRow(
		`INSERT INTO dishes 
			(name, description, image, video_link, calories, variation, day_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		dish.Name, dish.Description, dish.Image, dish.VideoLink, dish.Calories, dish.Variation, dish.DayTime,
	).Scan(&dish.ID)
}

func (r *DishRepo) FindByID(id int) (*models.Dish, error) {
	var dish *models.Dish
	err := r.db.Get(dish, `SELECT * FROM dishes WHERE id = $1`, id)
	return dish, err
}

func (r *DishRepo) GetAll(limit int, offset int) ([]models.Dish, error) {
	dish := []models.Dish{}
	if err := r.db.Select(&dish, `SELECT * FROM dishes LIMIT $1 OFFSET $2`, limit, offset); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return dish, nil
}

func (r *DishRepo) DeleteByID(id int) error {
	if err := r.db.QueryRow(`DELETE FROM dishes WHERE id = $1`, id).Scan(); err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *DishRepo) FindAllInCaloriesRange(gte float32, lte float32) ([]*models.Dish, error) {
	var dishes []*models.Dish

	if err := r.db.Select(&dishes, `SELECT * FROM dishes
		WHERE calories > $1 AND calories < $2
		ORDER BY NEWID()`, gte, lte); err != nil {
		return nil, err
	}

	return dishes, nil
}
