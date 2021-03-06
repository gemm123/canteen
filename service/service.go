package service

import (
	"github.com/gemm123/canteen/models"
	"github.com/gemm123/canteen/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository: repository}
}

type Service interface {
	FindAllProduct() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id int) error
	GetCash() (models.Money, error)
	UpdateCash(id int, moneyRequest models.Money) (models.Money, error)
	CreateUser(user models.User) (models.User, error)
	LoginUser(id int) (models.User, error)
}
