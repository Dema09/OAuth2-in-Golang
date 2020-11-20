package users

import (
	"OAuth2-In-Golang/config/database_config"
	"OAuth2-In-Golang/response/response_body"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	bcrypt2 "golang.org/x/crypto/bcrypt"
)

var(
	userCollection = database_config.ConnectDB().Database("user_dummy_database").Collection("user_data")
)

func (user *UserData) SaveUser() (interface{}, *response_body.Response){
	userPasswordInString := user.HashPassword()
	user.Password = userPasswordInString

	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil{
		return nil, response_body.NewStatusInternalServerError(fmt.Sprintf("Error when inserting new user data: %s", err))
	}
	return insertResult.InsertedID, nil
}

func (user *UserData) HashPassword() string{
	hash, _ := bcrypt2.GenerateFromPassword([]byte(user.Password), bcrypt2.DefaultCost)
	return string(hash)
}

func (user *UserData) FindUserDataByUsername(userLoginRequest LoginRequest)(*UserData, *response_body.Response){
	getUserData := userCollection.FindOne(context.TODO(),
		bson.M{
			"username": userLoginRequest.Username,
		},
		)
	if getUserData == nil{
		return nil, response_body.NewStatusNotFound(fmt.Sprintf("No data were matching with the username %s", userLoginRequest.Username))
	}

	userDecode := &UserData{}
	getUserData.Decode(userDecode)
	return userDecode, nil
}