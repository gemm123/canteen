package models

type User struct {
	UserID   uint
	Password string `gorm:"type:varchar(255)"`
}
