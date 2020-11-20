package service

import (
	"OAuth2-In-Golang/domain/users"
	"OAuth2-In-Golang/response/response_body"
	bcrypt2 "golang.org/x/crypto/bcrypt"
)

func RegisterUser(user users.UserData) (interface{}, *response_body.Response){
	userRegisterResponse, err := user.SaveUser()
	if err != nil{
		return "", err
	}
	return userRegisterResponse, nil
}

func LoginUser(user users.LoginRequest)(*response_body.Response, *response_body.Response){
	var currentUserData *users.UserData

	if err := user.Validate(); err != nil{
		return nil, err
	}

	loginResponse, err := currentUserData.FindUserDataByUsername(user)
	if err != nil{
		return nil, err
	}
	passwordInString := bcrypt2.CompareHashAndPassword([]byte(loginResponse.Password), []byte(user.Password))

	if loginResponse.Username != user.Username || passwordInString != nil{
		return nil, response_body.NewStatusUnauthorized("Your username or password is incorrect!")
	}
	return response_body.NewStatusOK(loginResponse), nil
}