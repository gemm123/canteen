package controller

import (
	"net/http"
	"strconv"

	"github.com/gemm123/canteen/models"
	"github.com/gin-gonic/gin"
)

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

func (ctr *controller) BuyProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := ctr.service.DeleteProduct(idInt)
	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/products")
}
