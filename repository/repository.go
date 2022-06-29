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
	FindProductByID(id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) error
	GetCash() (models.Money, error)
	GetCashByID(id int) (models.Money, error)
	UpdateCash(money models.Money) (models.Money, error)
}
