package repository

import "github.com/gemm123/canteen/models"

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}
