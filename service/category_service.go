package service

import (
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type CategoryService interface {
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id int64) error
	GetByID(id int64) (*model.Category, error)
	GetAll() ([]model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) Create(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) Update(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetByID(id int64) (*model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}