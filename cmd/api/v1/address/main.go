package main

import (
	_ "github.com/opas-keat/addressapi/cmd/api/v1/address/docs"

	"github.com/opas-keat/addressapi/cmd/api/v1/address/handlers"
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
	// database.Connect()
	// database.InitGeographie()
	// database.InitProvince()
	// database.InitAmphure()
	// database.InitDistrict()
	handlers.API()
}
