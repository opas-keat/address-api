package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opas-keat/addressapi/cmd/api/v1/address/model"
	"github.com/opas-keat/addressapi/database"
)

// InitGeographie is a function to create init geographie data to database
// @Summary create init geographie
// @Description create init geographie
// @Tags geographies
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseHTTP{}
// @Failure 503 {object} model.ResponseHTTP{}
// @Router /geographies/init [post]
func InitGeographie(c *fiber.Ctx) error {
	database.InitGeographie()
	return c.JSON(model.ResponseHTTP{
		Success: true,
		Message: "Success init geographie.",
	})
}
