package config

import (
	"fmt"
	"gorm-study/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	username = "root"
	password = "123456"
	host     = "127.0.0.1"
	port     = 3306
	name     = "go"
)

var DB *gorm.DB

func getDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC", username, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}

func init() {
	DB = getDb()
	DB.AutoMigrate(&entity.User{}, &entity.Role{})
}
