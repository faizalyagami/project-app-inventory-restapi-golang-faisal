package repository

import (
	"database/sql"
	"errors"
	"project-app-inventory-restapi-golang-faisal/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
}

type userRepository struct {
  	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	row := r.db.QueryRow(`SELECT id, username, password, role FROM users WHERE username=$1`, username)

	var user model.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	query := `INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`
	_, err := r.db.Exec(query, user.Username,user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}