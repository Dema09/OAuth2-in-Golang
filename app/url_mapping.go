package app

import "OAuth2-In-Golang/controller"

func UrlMapping(){

	router.POST("/v1/register", controller.RegisterUser)
	router.POST("/v1/login", controller.LoginUser)
}