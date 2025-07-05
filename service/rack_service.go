package service

import (
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type RackService interface {
	GetAll() ([]model.Rack, error)
	GetByID(id int64) (*model.Rack, error)
	Create(rack *model.Rack) error
	Update(rack *model.Rack) error
	Delete(id int64) error
}

type rackService struct {
	repo repository.RackRepository
}

func NewRackService(r repository.RackRepository) RackService {
	return &rackService{repo: r}
}

func (s *rackService) GetAll() ([]model.Rack, error) {
	return  s.repo.GetAll()
}

func (s *rackService) GetByID(id int64) (*model.Rack, error) {
	return s.repo.GetByID(id)
}

func (s rackService) Create(rack *model.Rack) error {
	return s.repo.Create(rack)
}

func (s *rackService) Update(rack *model.Rack) error {
	return s.repo.Update(rack)
}

func (s *rackService) Delete(id int64) error {
	return s.repo.Delete(id)
}