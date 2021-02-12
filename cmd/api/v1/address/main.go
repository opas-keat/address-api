package main

import (
	"github.com/opas-keat/addressapi/cmd/api/v1/address/handlers"
	"github.com/opas-keat/addressapi/database"
)

func main() {
	database.Connect()
	// database.InitGeographie()
	// database.InitProvince()
	// database.InitAmphure()
	// database.InitDistrict()
	handlers.API()

	// app := handlers.API()
	// err := app.Listen(":4000")

	// if err != nil {
	// 	panic(err)
	// }
}
