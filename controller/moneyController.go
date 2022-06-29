package controller

import (
	"log"
	"net/http"

	"github.com/gemm123/canteen/models"
	"github.com/gin-gonic/gin"
)

func (ctr *controller) UpdateCash(c *gin.Context) {
	cash := c.PostForm("cash")

	moneyRequest := models.Money{
		Cash: cash,
	}

	log.Println(moneyRequest.Cash)

	_, err := ctr.service.UpdateCash(1, moneyRequest)
	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/products")
}
