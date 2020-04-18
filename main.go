package main

import (
	"github.com/Tez-Private/jetbook/config"
	"github.com/Tez-Private/jetbook/db"
	"github.com/Tez-Private/jetbook/model"
	"github.com/Tez-Private/jetbook/route"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
)

func main() {
	config.ReadLocalEnvConfig()
	mysql := db.SetupDB()
	defer mysql.Close()

	mysql.AutoMigrate(
		&model.Users{},
		&model.Rents{},
		&model.Books{}, //Need    ALTER TABLE books CONVERT TO CHARACTER SET utf8mb4;
	)
	/*
		mysql.Model(&model.Rents{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")
		mysql.Model(&model.Rents{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	*/

	api := gin.New()
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	route.V1(api)
	api.Run()
}
