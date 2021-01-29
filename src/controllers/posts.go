package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CreatePost(w http.ResponseWriter, r *http.Request){
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w,http.StatusUnprocessableEntity,err)
		return
	}
	var post models.Posts
	if err := json.Unmarshal(bodyRequest,&post);err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	ID, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	post.AuthorID = uint(ID)
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	postRepo := repositories.NewRepositoryPosts(db)
	if err := post.Treatment(); err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	post, err = postRepo.Create(post)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusOK,post)
}
func FindPost(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.GetBD()
	if err != err {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	var post models.Posts
	postRepo := repositories.NewRepositoryPosts(db)
	post, err = postRepo.Find(uint(ID))
	if err != nil{
		response.Error(w,http.StatusUnprocessableEntity,err)
		return
	}
	response.JSON(w,http.StatusOK,post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	IDPost,err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	IDToken, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	db,err := database.GetBD()
	if err != nil {
		response.Error(w, http.StatusInternalServerError,err)
		return
	}
	postRepo := repositories.NewRepositoryPosts(db)
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil{
		response.Error(w,http.StatusUnauthorized,err)
		return
	}
	var posts models.Posts
	if err = json.Unmarshal(bodyRequest,&posts);err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if err = postRepo.Update(posts,uint(IDPost),uint(IDToken)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func ListPost(w http.ResponseWriter, r *http.Request){
	findParam := strings.ToLower(r.URL.Query().Get("title"))
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	postRepo := repositories.NewRepositoryPosts(db)
	var posts []models.Posts
	posts = postRepo.List(findParam)
	response.JSON(w,http.StatusOK,posts)
}
func DeletePost(w http.ResponseWriter, r *http.Request){
	IDtoken, err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	parameters := mux.Vars(r)
	IDPost, err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w, http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	postRepo := repositories.NewRepositoryPosts(db)
	var post models.Posts
	post,err = postRepo.Find(uint(IDPost))
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	if post.AuthorID != uint(IDtoken) {
		response.Error(w,http.StatusForbidden,errors.New("Você não possui autorização para excluir esse post"))
		return
	}
	if err = postRepo.Delete(uint(IDPost)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)

}
func Like(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	IDPost,err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	ID,err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	PostRepo := repositories.NewRepositoryPosts(db)
	if err = PostRepo.Like(uint(IDPost),uint(ID)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}
func UnLike(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	IDPost,err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	db, err := database.GetBD()
	if err != nil {
		response.Error(w,http.StatusInternalServerError,err)
		return
	}
	ID,err := authentication.TakeID(r)
	if err != nil {
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	PostRepo := repositories.NewRepositoryPosts(db)
	if err = PostRepo.UnLike(uint(IDPost),uint(ID)); err != nil{
		response.Error(w,http.StatusBadRequest,err)
		return
	}
	response.JSON(w,http.StatusNoContent,nil)
}