package controller

import (
	"OAuth2-In-Golang/domain/users"
	"OAuth2-In-Golang/response/response_body"
	"OAuth2-In-Golang/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context){
	var user users.UserData
	if err := c.ShouldBindJSON(&user); err != nil{
		responseBody := response_body.NewStatusBadRequest("Invalid JSON Body")
		c.JSON(responseBody.StatusCode, responseBody)
		return
	}

	registerUserResponse, err := service.RegisterUser(user)
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, registerUserResponse)
}

func LoginUser(c *gin.Context){
	var userLoginRequest users.LoginRequest
	if err := c.ShouldBindJSON(&userLoginRequest); err != nil{
		responseBody := response_body.NewStatusBadRequest("Invalid JSON Body!")
		c.JSON(responseBody.StatusCode, responseBody)
		return
	}

	loginResponse, err := service.LoginUser(userLoginRequest)
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(loginResponse.StatusCode, loginResponse)
}