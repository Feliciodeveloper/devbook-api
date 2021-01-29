package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(ID uint)(string,error){
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,permissions)
	return token.SignedString([]byte(config.SecretKey))
}
func ValidateToken(r *http.Request)error{
	tokenString := takeToken(r)
	token, err := jwt.Parse(tokenString,returnKey)
	if err != nil {
		return err
	}
	if _,ok :=token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token invalido")
}
func takeToken(r *http.Request)string{
	token := r.Header.Get("Authorization")
	if len(strings.Split(token," ")) == 2{
		return strings.Split(token," ")[1]
	}
	return ""
}
func TakeID(r *http.Request)(uint64, error){
	tokenString := takeToken(r)
	token, err := jwt.Parse(tokenString,returnKey)
	if err != nil {
		return 0,err
	}
	if toAllow, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f",toAllow["id"]),10,64)
		if err != nil{
			return 0,err
		}
		return userID, nil
	}
	return 0, errors.New("Token invalido")
}
func returnKey(token *jwt.Token)(interface{},error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
