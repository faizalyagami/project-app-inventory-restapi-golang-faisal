package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-faisal/model"
)

type WarehouseRepository interface {
	GetAll() ([]model.Warehouse, error)
	GetByID(id int64) (*model.Warehouse, error)
	Create(warehouse *model.Warehouse) error
	Update(warehouse *model.Warehouse) error
	Delete(id int64) error
}

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) GetAll()([]model.Warehouse, error) {
	rows, err := r.db.Query(`SELECT * FROM warehouses ORDER BY id`)
	if err != nil {
		return  nil, err
	}
	defer rows.Close()

	var warehouses []model.Warehouse
	for rows.Next() {
		var w model.Warehouse
		if err := rows.Scan(&w.ID, &w.Name); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return  warehouses, nil
}

func (r *warehouseRepository) GetByID(id int64) (*model.Warehouse, error) {
	row := r.db.QueryRow(`SELECT * FROM warehouses WHERE id=$1`, id)

	var w model.Warehouse
	if err := row.Scan(&w.ID, &w.Name); err != nil {
		return  nil, err
	}
	return &w, nil
}

func (r *warehouseRepository) Create(warehouse *model.Warehouse) error {
	_, err := r.db.Exec(`INSERT INTO warehouses(name) VALUES($1) RETURNING id`, warehouse.Name)
	return  err
}
func (r *warehouseRepository) Update(warehouse *model.Warehouse) error {
	_, err := r.db.Exec(`UPDATE FROM warehouses SET name=$1`, warehouse.Name)
	return  err
}
func (r *warehouseRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM warehouses WHERE id=$1`, id)
	return  err
}