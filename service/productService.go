package service

import "github.com/gemm123/canteen/models"

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
