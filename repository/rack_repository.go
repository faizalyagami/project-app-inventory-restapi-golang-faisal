package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-faisal/model"
)

type RackRepository interface {
	GetAll() ([]model.Rack, error)
	GetByID(id int64) (*model.Rack, error)
	Create(rack *model.Rack) error
	Update(rack *model.Rack) error
	Delete(id int64) error
}

type rackRepository struct {
	db *sql.DB
}

func NewRackRepository(db *sql.DB) RackRepository {
	return &rackRepository{db: db}
}

func (r *rackRepository) GetAll()([]model.Rack, error) {
	rows, err := r.db.Query(`SELECT id, name FROM racks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var racks []model.Rack
	for rows.Next(){
		var r model.Rack
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		racks = append(racks, r)
	}
	return racks, nil
}
func (r *rackRepository) GetByID(id int64)(*model.Rack, error) {
	row := r.db.QueryRow(`SELECT id, name FROM racks WHERE id=$1`, id)

	var rack model.Rack
	if err := row.Scan(&rack.ID, &rack.Name); err !=nil {
		return nil, err
	}
	return &rack, nil
}


func (r *rackRepository) Create(rack *model.Rack) error {
	_, err := r.db.Exec(`INSERT INTO racks (name) VALUES($1) RETURNING id`, rack.Name)
	return err
}
func (r *rackRepository) Update(rack *model.Rack) error {
	_, err := r.db.Exec(`UPDATE racks SET name=$1 WHERE id=$2`, rack.Name, rack.ID)
	return err
}
func (r *rackRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM racks WHERE id=$1`, id)
	return  err
}