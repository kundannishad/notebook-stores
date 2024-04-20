package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const (
	username = "root"
	password = "Password@123"
	hostname = "127.0.0.1:3306"
	database = "golandDB"
)

func ConnectionString() string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, database)
}

func DbConnect() {

	fmt.Println(ConnectionString())
	dbs, err := gorm.Open("mysql", ConnectionString()+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println("Errot", err)
		panic(err)
	}
	db = dbs
}

func GetDb() *gorm.DB {
	return db
}
