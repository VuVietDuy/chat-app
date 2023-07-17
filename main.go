package main

import (
	"kafka/entity"
	"kafka/pkg/mysql"
)

func main() {
	db := mysql.New(mysql.Config{
		Username: "root",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Database: "demo",
	})

	db.DB.AutoMigrate(&entity.User{})
}
