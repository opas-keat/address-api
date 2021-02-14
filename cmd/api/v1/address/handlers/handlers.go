package handlers

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/opas-keat/addressapi/cmd/api/v1/address/service"
	"github.com/spf13/viper"
)

// API constructs an http.Handler with all application routes defined.
func API() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	var port = viper.GetString("app.port")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/docs/*", swagger.Handler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🐣 Address API.",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})

	v1.Post("/geographies/init", service.InitGeographie)

	v1.Post("/provinces/init", service.InitProvince)
	v1.Get("/provinces", service.GetProvinces)
	v1.Get("/provinces/name/th", service.SearchProvincesByName)

	v1.Post("/amphures/init", service.InitAmphure)
	v1.Get("/amphures", service.GetAmphures)
	v1.Get("/amphures/province/:province_id", service.GetAmphuresByProvinceID)

	v1.Post("/districts/init", service.InitDistrict)
	v1.Get("/districts", service.GetDistricts)
	v1.Get("/districts/amphure/:amphure_id", service.GetDistrictsByAmphureID)

	err = app.Listen(port)
	if err != nil {
		panic(err)
	}
}
