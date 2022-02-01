package database

import (
	"KanbanBoardAPI/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host = "localhost"
	user = "postgres"
	password = "amaterasu"
	dbname = "kanban_db"
	port = "5432"

	db *gorm.DB
	err error
)


func StartDB(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database =>", err.Error())
		return
	}

	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("error migration =>", err.Error())
		return
	}

	fmt.Println("Database Connected")
}

func GetDB() *gorm.DB {
	return db
}
