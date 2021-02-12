package main

import (
	"github.com/opas-keat/addressapi/database"
)

func main() {
	database.Connect()
	// fmt.Println(addressapi.Config.User)
}
