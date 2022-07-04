package config

import (
	"log"

	"github.com/gemm123/canteen/models"
	"github.com/gemm123/canteen/seeds"
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
	if err := seeds.All().Run(DB); err != nil {
		log.Fatal("failerd seed to database")
	}
}

func CloseDB() {
	DBSQL, _ := DB.DB()
	DBSQL.Close()
}
