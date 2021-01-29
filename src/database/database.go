package database

import (
	"api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetBD()(*gorm.DB,error) {
	db, err := gorm.Open(mysql.Open(config.StringConn),&gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	return db,nil
}