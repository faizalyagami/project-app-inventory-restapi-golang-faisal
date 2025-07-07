package service

import (
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type ItemService interface {
  GetAll() ([]model.Item, error)
  GetByID(id int64) (*model.Item, error)
  Create(item *model.Item) error
  Update(item *model.Item) error
  Delete(id int64) error
  GetLowStockItems(threshold int64) ([]model.Item, error)
  GetPaginated(page, limit int) ([]model.Item, error)
}

type itemService struct {
  repo repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) ItemService {
	return &itemService{repo: r}
}

func (s *itemService) GetAll() ([]model.Item, error) {
	return s.repo.GetAll()
}

func (s *itemService) GetByID(id int64) (*model.Item, error) {
	return  s.repo.GetByID(id)
}

func (s *itemService) Create(item *model.Item) error {
	return  s.repo.Create(item)
}

func (s *itemService) Update(item *model.Item) error {
	return s.repo.Update(item)
}

func (s *itemService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *itemService) GetLowStockItems(thresold int64) ([]model.Item,error) {
	return s.repo.GetLowStockItems(thresold)
}

func (s *itemService) GetPaginated(page, limit int) ([]model.Item, error) {
	offset := (page - 1) *limit
	return s.repo.GetPaginated(offset, limit)
}