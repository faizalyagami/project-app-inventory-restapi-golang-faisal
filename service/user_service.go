package service

import (
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userService struct {
  repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}
func (s *userService) GetByID(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}
func (s *userService) GetByUsername(username string) (*model.User, error) {
	return s.repo.GetByUsername(username)
}
func (s *userService) Create(user *model.User) (*model.User, error) {
	return s.repo.Create(user)
}
func (s *userService) Update(user *model.User) error {
	return s.repo.Update(user)
}
func (s *userService) Delete(id int64) error {
	return s.repo.Delete(id)
}
