package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/safety"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CreateUSer(w http.ResponseWriter, r *http.Request){
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w,http.StatusUnprocessableEntity,err)
		return
	}
	var user models.Users
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if err = user.Treatment("create"); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	usersRepo := repositories.NewRepositoryUsers(db)
	user = usersRepo.Create(user)
	response.JSON(w,http.StatusCreated,user)
}
func ListUsers(w http.ResponseWriter, r *http.Request){
	findParam := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	users := userRepo.List(findParam)
	response.JSON(w,http.StatusOK,users)
}
func FindUser(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	var user models.Users
	user = userRepo.Find(ID)
	response.JSON(w,http.StatusOK,user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	UserIDToken, err := authentication.TakeID(r)
	if ID != UserIDToken {
		response.Error(w,http.StatusForbidden,errors.New("Sem autorização para realizar essa ação"))
		return
	}
	if err != nil {
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	var user models.Users
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w,http.StatusUnprocessableEntity,err)
		return
	}
	if err = json.Unmarshal(bodyRequest,&user); err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if err := user.Treatment("update"); err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	user = userRepo.Update(ID,user)
	response.JSON(w,http.StatusOK,user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	UserIDToken, err := authentication.TakeID(r)
	if ID != UserIDToken {
		response.Error(w,http.StatusForbidden,errors.New("Sem autorização para realizar essa ação"))
		return
	}
	if err != nil {
		response.Error(w,http.StatusBadRequest, err)
		return
	}
	db, err := database.GetBD()
	if err != nil{
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	if err = userRepo.Delete(ID); err != nil {
		response.Error(w, http.StatusUnprocessableEntity,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func UpdatePassword(w http.ResponseWriter, r *http.Request){
	IDToken, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if ID != IDToken {
		response.Error(w, http.StatusForbidden, errors.New("Não é possivel alterar a senha de outro usuario"))
		return
	}
	bodyRequest, err := ioutil.ReadAll(r.Body)
	var password models.Password
	if err := json.Unmarshal(bodyRequest,&password); err != nil {
		response.Error(w, http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(db)
	PassRepo,err := userRepo.FindPassword(ID)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if err = safety.VerifyPassword(PassRepo,password.Old); err != nil {
		response.Error(w,http.StatusUnauthorized,errors.New("Senha de verificação difere do banco"))
		return
	}
	passwordHash,err := safety.Hash(password.New)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := userRepo.UpdatePassword(string(passwordHash),ID); err != nil{
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func AddFriend(w http.ResponseWriter, r *http.Request){
	ID, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	parameters := mux.Vars(r)
	IDFriend, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if ID == IDFriend {
		response.Error(w, http.StatusForbidden,errors.New("Não e possivel adicionar você mesmo como amigo"))
		return
	}
	bd, err := database.GetBD()
	if err != nil {
		response.Error(w, http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(bd)
	if err = userRepo.AddFriend(uint(ID),uint(IDFriend)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func RemoveFriend(w http.ResponseWriter, r *http.Request){
	ID, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	parameters := mux.Vars(r)
	IDFriend, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if ID == IDFriend {
		response.Error(w, http.StatusForbidden,errors.New("Não e possivel adicionar você mesmo como amigo"))
	}
	bd, err := database.GetBD()
	if err != nil {
		response.Error(w, http.StatusInternalServerError,err)
		return
	}
	userRepo := repositories.NewRepositoryUsers(bd)
	if err = userRepo.RemoveFriend(uint(ID),uint(IDFriend)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func ListFriends(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	var users []models.Users
	usersRepo := repositories.NewRepositoryUsers(db)
	users,err = usersRepo.ListFriends(uint(ID))
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusOK,users)
}