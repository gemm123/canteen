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
}
