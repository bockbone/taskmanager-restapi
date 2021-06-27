package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bockbone/taskmanager/common"
)

//Handler for POST - /users/register
//Register user
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource

	//Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User Data",
			500,
		)
		return
	}

	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	//Insert document
	repo.CreateUser(user)


	//Cleanup the hash password to eliminate it from response
	user.HashPassword = nil
	if j,err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}