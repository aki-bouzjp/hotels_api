package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

func New() *DB {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	url := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
	}
	return &DB{
		Conn: conn,
	}
}
