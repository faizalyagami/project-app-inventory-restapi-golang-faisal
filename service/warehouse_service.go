package service

import (
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type WarehouseService interface {
	GetAll() ([]model.Warehouse, error)
	GetByID(id int64) (*model.Warehouse, error)
	Create(warehouse *model.Warehouse) error
	Update(warehouse *model.Warehouse) error
	Delete(id int64) error
}

type warehouseService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseService(r repository.WarehouseRepository) WarehouseService {
	return &warehouseService{repo: r}
}

func (s *warehouseService) GetAll() ([]model.Warehouse, error) {
	return s.repo.GetAll()
}
func (s *warehouseService) GetByID(id int64) (*model.Warehouse, error) {
	return s.repo.GetByID(id)
}
func (s *warehouseService) Create(warehouse *model.Warehouse) error {
	return s.repo.Create(warehouse)
}
func (s *warehouseService) Update(warehouse *model.Warehouse) error {
	return s.repo.Update(warehouse)
}
func (s *warehouseService) Delete(id int64) error {
	return s.repo.Delete(id)
}