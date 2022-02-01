package main

import (
	"KanbanBoardAPI/database"
	"KanbanBoardAPI/router"
	"log"
)

func main(){

	// connecting database
	database.StartDB()

	app := router.Router()

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("error start server =>", err.Error())
		return
	}

}
