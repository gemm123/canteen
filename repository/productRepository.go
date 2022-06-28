package repository

import (
	"github.com/gemm123/canteen/helper"
	"github.com/gemm123/canteen/models"
)

func (r *repository) FindAllProduct() ([]models.Product, error) {
	var products []models.Product
	query := `select prod.id, prod.updated_at, prod.name, prod.image, prod.desc, prod.price from products prod`
	rows, err := r.DB.Raw(query).Rows()
	if err != nil {
		return products, err
	}

	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.UpdatedAt, &product.Name, &product.Image, &product.Desc, &product.Price)
		if err != nil {
			return products, err
		}
		product.Time = helper.FormatDate(product.UpdatedAt)
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindProductByID(id int) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.DB.Create(&product).Error
	return product, err
}

func (r *repository) DeleteProduct(product models.Product) error {
	err := r.DB.Unscoped().Delete(&product).Error
	return err
}
