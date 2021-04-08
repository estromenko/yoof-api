package repo

import (
	"database/sql"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type DishRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DishRepo {
	return &DishRepo{
		db: db,
	}
}

func (r *DishRepo) FindByID(id int) (*models.Dish, error) {
	var dish *models.Dish
	err := r.db.Get(dish, `SELECT * FROM dishes WHERE id = $1`, id)
	return dish, err
}

func (r *DishRepo) GetAll() ([]*models.Dish, error) {
	var dish []*models.Dish
	if err := r.db.Select(dish, `SELECT * FROM dishes`); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return dish, nil
}
