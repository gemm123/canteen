package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gemm123/canteen/helper"
	"github.com/gemm123/canteen/models"
	"github.com/gin-gonic/gin"
)

func (ctr *controller) PostUser(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")

	splitID := strings.Split(id, "")
	if len(splitID) != 5 {
		return
	}

	sumSliceString := splitID[0:3]
	sumSliceInt := helper.ConvertSliceStringToInteger(sumSliceString)
	sum := helper.Sum(sumSliceInt)

	resultString := splitID[3] + splitID[4]
	result, _ := strconv.Atoi(resultString)

	if sum != result {
		return
	}

	passwordHashed, err := helper.HashPassword(password)
	if err != nil {
		return
	}

	userID, _ := strconv.Atoi(id)
	newUser := models.User{
		UserID:   uint(userID),
		Password: passwordHashed,
	}

	ctr.service.CreateUser(newUser)

	c.Redirect(http.StatusFound, "/login")
}