package controller

import (
	"net/http"

	"github.com/gemm123/canteen/models"
	"github.com/gin-gonic/gin"
)

func (ctr *controller) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func (ctr *controller) Products(c *gin.Context) {
	products, err := ctr.service.FindAllProduct()
	if err != nil {
		c.HTML(http.StatusBadRequest, "products.tmpl", gin.H{
			"data": err,
		})
		return
	}

	c.HTML(http.StatusOK, "products.tmpl", gin.H{
		"data": products,
	})
}

func (ctr *controller) CreateProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "create-product.tmpl", gin.H{})
}

func (ctr *controller) PostCreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	price := c.PostForm("price")

	file, err := c.FormFile("image")
	if err != nil {
		return
	}

	path := "static/images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		return
	}

	product := models.Product{
		Name:  name,
		Desc:  desc,
		Price: price,
		Image: path,
	}

	_, err = ctr.service.CreateProduct(product)
	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/products")
}
