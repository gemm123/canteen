package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gemm123/canteen/helper"
	"github.com/gemm123/canteen/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (ctr *controller) PostRegister(c *gin.Context) {
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

	_, err = ctr.service.CreateUser(newUser)
	if err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func (ctr *controller) PostLogin(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")

	idString, _ := strconv.Atoi(id)
	user, err := ctr.service.LoginUser(idString)
	if err != nil {
		return
	}

	ok := helper.CheckPasswordHash(password, user.Password)
	if !ok {
		return
	}

	session := sessions.Default(c)
	session.Set("id", id)
	err = session.Save()
	if err != nil {
		return
	}
	authenticated = 1

	c.Redirect(http.StatusFound, "/")
}

func (ctr *controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("id")
	session.Save()
	authenticated = 0

	c.Redirect(http.StatusFound, "/")
}
