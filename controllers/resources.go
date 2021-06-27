package controllers

import "github.com/bockbone/taskmanager/models"

type (
	//For POST - /user/register
	UserResource struct {
		Data	models.User		`json:"data"`	
	}

	//For POST - /user/login
	LoginResource struct {
		Data	LoginModel		`json:"data"`	
	}

	//Response for authorized user POST - /user/login
	AuthUserResource struct {
		Data AuthUserModel	`json:"data"`
	}

	//Model for authentication
	LoginModel struct {
		Email	string	`json:"email"`
		Password	string	`json:"password"`
	}

	//Model for authorizes user with access token
	AuthUserModel struct {
		User models.User `json:"user"`
		Token string	`json:"token"`
	}
)