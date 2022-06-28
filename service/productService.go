package service

import (
	"os"

	"github.com/gemm123/canteen/models"
)

func (s *service) FindAllProduct() ([]models.Product, error) {
	products, err := s.repository.FindAllProduct()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) CreateProduct(product models.Product) (models.Product, error) {
	product, err := s.repository.CreateProduct(product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) DeleteProduct(id int) error {
	product, err := s.repository.FindProductByID(id)
	if err != nil {
		return err
	}

	path := product.Image
	err = os.Remove(path)
	if err != nil {
		return err
	}

	err = s.repository.DeleteProduct(product)
	if err != nil {
		return err
	}

	return nil
}
