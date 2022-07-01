package service

import (
	"github.com/gemm123/canteen/models"
)

func (s *service) CreateUser(user models.User) (models.User, error) {
	user, err := s.repository.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) LoginUser(id int) (models.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
