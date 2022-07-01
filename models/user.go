package models

type User struct {
	UserID   uint   `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`
}
