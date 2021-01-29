package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConn = ""
	Port = 0
	SecretKey []byte
)

func Loading(){
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}