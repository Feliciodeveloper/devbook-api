package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/safety"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w,http.StatusUnprocessableEntity, err)
		return
	}
	var user models.Users
	if err = json.Unmarshal(bodyRequest,&user);err != nil {
		response.Error(w,http.StatusUnprocessableEntity,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	UserDB,err := userRepo.Login(user.Email)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if err = safety.VerifyPassword(UserDB.Password,user.Password); err != nil{
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	token, err := authentication.GenerateToken(UserDB.ID)
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	w.Write([]byte(token))
}
