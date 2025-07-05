package service

import (
	"errors"
	"project-app-inventory-restapi-golang-faisal/model"
	"project-app-inventory-restapi-golang-faisal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (*model.User, error)
	Register(username,email, password string) (*model.User, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return  &authService{repo: r}
}

func (s *authService) Login(username, password string) (*model.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	user.Password = ""
	return user, nil
}

func (s *authService) Register(username, email, hashedPassword string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Email: email,
		Password: hashedPassword,
		Role: "staff",
	}
	user, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}