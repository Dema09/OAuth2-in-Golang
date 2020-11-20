package users

import "OAuth2-In-Golang/response/response_body"

type LoginRequest struct{
	Username string `json: "username"`
	Password string `json: "password"`
}

func (loginRequest *LoginRequest) Validate() *response_body.Response{
	if loginRequest.Username == "" || loginRequest.Password == ""{
		return response_body.NewStatusBadRequest("Your username or password is empty!")
	}
	return nil
}