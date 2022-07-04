package seeds

import (
	"github.com/gemm123/canteen/seed"
	"gorm.io/gorm"
)

func All() seed.Seed {
	return seed.Seed{
		Name: "Initial money",
		Run: func(db *gorm.DB) error {
			CreateMoney(db, "Rp. 0")
			return nil
		},
	}
}
