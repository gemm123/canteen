package seeds

import (
	"github.com/gemm123/canteen/models"
	"gorm.io/gorm"
)

func CreateMoney(db *gorm.DB, cash string) error {
	return db.Create(&models.Money{Cash: cash}).Error
}
