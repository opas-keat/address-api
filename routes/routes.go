package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"omsoft.com/addressapi/services"
)

//SetupRoutes for ADDRESS API
func SetupRoutes() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/docs/*", swagger.Handler)

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})

	v1.Get("/provinces", services.GetProvinces)

	v1.Get("/amphures", services.GetAmphures)
	v1.Get("/amphures/province/:province_id", services.GetAmphuresByProvinceID)

	v1.Get("/districts", services.GetDistricts)
	v1.Get("/districts/amphure/:amphure_id", services.GetDistrictsByAmphureID)

	return app

}
