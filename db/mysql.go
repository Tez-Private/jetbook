package db

import (
	"log"

	config "github.com/Tez-Private/jetbook/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//var DB *gorm.DB
var Mysql *gorm.DB

func getDbConfig() string {
	user := config.EnvConfig.MYSQL_USER
	pass := config.EnvConfig.MYSQL_PASS
	dbname := config.EnvConfig.MYSQL_DBNAME
	host := config.EnvConfig.MYSQL_HOST
	access := config.EnvConfig.MYSQL_ACCESS

	// LocalでのMySQL接続を想定して記述 AWS or GCPを使う場合にはもしかすると修正が必要かもしれない
	connect := user + ":" + pass + "@" + access + "/" + dbname + "?charset=utf8&parseTime=True&loc=" + host
	return connect
}

func SetupDB() *gorm.DB {
	db, err := gorm.Open("mysql", getDbConfig())
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	Mysql = db
	return Mysql
}
