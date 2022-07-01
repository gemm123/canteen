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
	session := sessions.Default(c)

	splitID := strings.Split(id, "")
	if len(splitID) != 5 {
		session.Set("err", "ID must 5 character number")
		session.Save()
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sumSliceString := splitID[0:3]
	sumSliceInt := helper.ConvertSliceStringToInteger(sumSliceString)
	sum := helper.Sum(sumSliceInt)

	resultString := splitID[3] + splitID[4]
	result, _ := strconv.Atoi(resultString)

	if sum != result {
		session.Set("err", "2 last digits must the sum of 3 first digits")
		session.Save()
		c.Redirect(http.StatusFound, "/register")
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
	session := sessions.Default(c)

	idString, _ := strconv.Atoi(id)
	user, err := ctr.service.LoginUser(idString)
	if err != nil {
		session.Set("err", "wrong id or password")
		session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}

	ok := helper.CheckPasswordHash(password, user.Password)
	if !ok {
		session.Set("err", "wrong id or password")
		session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}

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
