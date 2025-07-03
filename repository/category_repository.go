package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-faisal/model"
)

type CategoryRepository interface {
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id int64) error
	GetByID(id int64) (*model.Category, error)
	GetAll() ([]model.Category, error)
}

type categoryRepo struct {
  db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) Create(category *model.Category) error {
	err := r.db.QueryRow(`INSERT INTO categories(name) VALUES($1) RETURNING id`, category.Name).Scan(&category.ID)
	return err
}

func (r *categoryRepo) Update(category *model.Category) error {
	_, err := r.db.Exec(`UPDATE categories SET name=$1 WHERE id=$2`, category.ID, category.Name)
	return err
}

func (r *categoryRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM categories WHERE id=$1`, id)
	return err
}

func (r *categoryRepo) GetByID(id int64) (*model.Category, error) {
	row := r.db.QueryRow(`SELECT id, name FROM categories WHERE id=$1`, id)
	var category model.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return  nil, err
	}
	return  &category, nil
}

func (r *categoryRepo) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query(`SELECT id, name FROM categories ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return  nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}