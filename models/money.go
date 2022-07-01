package models

type Money struct {
	ID   uint
	Cash string `gorm:"type:varchar(255);default:Rp. 0"`
}
