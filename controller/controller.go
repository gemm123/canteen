package controller

import (
	"log"
	"net/http"

	"github.com/gemm123/canteen/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var authenticated int

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{service: service}
}

func (ctr *controller) Home(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"id":            id,
		"authenticated": authenticated,
	})
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

	session := sessions.Default(c)
	id := session.Get("id")

	c.HTML(http.StatusOK, "products.tmpl", gin.H{
		"data":          products,
		"money":         money,
		"id":            id,
		"authenticated": authenticated,
	})
}

func (ctr *controller) CreateProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "create-product.tmpl", gin.H{})
}

func (ctr *controller) Login(c *gin.Context) {
	session := sessions.Default(c)
	err := session.Get("err")
	session.Delete("err")
	session.Save()
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"err": err,
	})
}

func (ctr *controller) Register(c *gin.Context) {
	session := sessions.Default(c)
	err := session.Get("err")
	session.Delete("err")
	session.Save()
	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"err": err,
	})
}
