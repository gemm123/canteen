package config

import (
	"log"

	"github.com/gemm123/canteen/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConfigDB(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("err")
	}

	log.Println("Connected to database")

}

func MigrateDB() {
	DB.AutoMigrate(&models.Product{}, &models.Money{}, &models.User{})
}

func CloseDB() {
	DBSQL, _ := DB.DB()
	DBSQL.Close()
}
