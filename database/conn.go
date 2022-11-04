package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	db *gorm.DB
}
var Database dbInstance

func SqliteDBC () {
	db, err := gorm.Open(sqlite.Open("database/db/db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully !")
	db.Logger = logger.Default.LogMode(logger.Info)
	Database = dbInstance{db: db}	
}
