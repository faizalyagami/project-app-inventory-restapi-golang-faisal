package repository

import (
	"database/sql"
	"errors"
	"project-app-inventory-restapi-golang-faisal/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userRepository struct {
  	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, username, email role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err:= rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetByID(id int64) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, role, FORM users WHERE id=$1", id)
	var user model.User
	if err:= row.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
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

func (r *userRepository) Update(user *model.User) error {
	query := `UPDATE users SET username=$1, email=$2, role=$3, WHERE id=$4`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Role, user.ID)
	return err
}

func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return  err
}