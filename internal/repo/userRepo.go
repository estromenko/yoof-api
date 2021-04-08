package repo

import (
	"database/sql"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) FindByID(id int) (*models.User, error) {
	var user models.User
	if err := u.db.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Get(&user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Create(user *models.User) error {
	if _, err := u.FindByEmail(user.Email); err != nil && err != sql.ErrNoRows {
		return err
	}
	return u.db.QueryRow(`INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id`,
		&user.Email,
		&user.Username,
		&user.Password,
	).Scan(&user.ID)
}

func (u *UserRepo) Update(user *models.User) error {
	if _, err := u.FindByEmail(user.Email); err != nil && err != sql.ErrNoRows {
		return err
	}

	return u.db.QueryRow(`UPDATE users SET username = $1, password = $2`,
		&user.Username,
		&user.Password,
	).Scan()
}
