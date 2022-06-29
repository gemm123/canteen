package controller

import (
	"log"
	"net/http"

	"github.com/gemm123/canteen/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{service: service}
}

func (ctr *controller) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func (ctr *controller) Products(c *gin.Context) {
	products, err := ctr.service.FindAllProduct()
	if err != nil {
		log.Println(err)
		return
	}

	money, err := ctr.service.GetCash()
	if err != nil {
		log.Println(err)
		return
	}

	c.HTML(http.StatusOK, "products.tmpl", gin.H{
		"data":  products,
		"money": money,
	})
}

func (ctr *controller) CreateProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "create-product.tmpl", gin.H{})
}
