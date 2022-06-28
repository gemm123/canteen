package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Image string `gorm:"type:varchar(255)"`
	Desc  string `gorm:"type:varchar(255)"`
	Price string `gorm:"type:varchar(255)"`
	Time  string `gorm:"-:all"`
}
