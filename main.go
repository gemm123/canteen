package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gemm123/canteen/config"
	"github.com/gemm123/canteen/controller"
	"github.com/gemm123/canteen/render"
	"github.com/gemm123/canteen/repository"
	"github.com/gemm123/canteen/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cant get .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	config.ConfigDB(dsn)
	config.MigrateDB()
	defer config.CloseDB()

	r := gin.Default()
	r.HTMLRender = render.LoadTemplates("./templates")
	r.Static("/static", "./static")

	repository := repository.NewRepository(config.DB)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	r.GET("/", controller.Home)
	r.GET("/products", controller.Products)
	r.GET("/create-product", controller.CreateProduct)
	r.POST("/create-product", controller.PostCreateProduct)

	r.Run()
}
