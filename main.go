package main

import (
	"log"

	_ "omsoft.com/addressapi/docs"

	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/routes"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @query.collection.format multi
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
