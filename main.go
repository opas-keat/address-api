package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/models"
)

// func setupRoutes(app *fiber.App) {
// 	app.Get("/api/v1/provinces", province.GetProvinces)
// }

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=P@ssw0rd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// database.DBConn.AutoMigrate(&province.Province{})
	// fmt.Println("Database Migrated")
}

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	initDatabase()
	// setupRoutes(app)

	app.Get("/api/v1/provinces", models.GetProvinces)

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}

	// log.Fatal(app.Listen(":3000"))
}
