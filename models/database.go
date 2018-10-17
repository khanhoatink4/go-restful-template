package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	// db, err := gorm.Open("sqlite3", "db.db")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=hoank dbname=go password=hoank sslmode=disable")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	Migrate(db)
	DB = db
	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&ArticleModel{})
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
