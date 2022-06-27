package repository

import (
	"github.com/gemm123/canteen/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB: DB}
}

type Repository interface {
	FindAllProduct() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}
