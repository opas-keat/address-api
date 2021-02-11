package main

import (
	"log"

	_ "omsoft.com/addressapi/docs"

	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/routes"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	app := routes.SetupRoutes()
	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
