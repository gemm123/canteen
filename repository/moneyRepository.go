package repository

import "github.com/gemm123/canteen/models"

func (r *repository) GetCash() (models.Money, error) {
	var money models.Money
	err := r.DB.First(&money).Error
	return money, err
}

func (r *repository) GetCashByID(id int) (models.Money, error) {
	var money models.Money
	err := r.DB.First(&money, id).Error
	return money, err
}

func (r *repository) UpdateCash(money models.Money) (models.Money, error) {
	err := r.DB.Save(&money).Error
	return money, err
}
