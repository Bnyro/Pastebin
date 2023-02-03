package database

import (
	"flag"

	"github.com/pastebin/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

const DB_URI = "postgres://postgres@localhost/pastebin"

func Init() {
	dbUri := flag.String("db", DB_URI, "database uri")
	flag.Parse()

	var err error
	Db, err = gorm.Open(postgres.Open(*dbUri))
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&entities.Entry{})
}