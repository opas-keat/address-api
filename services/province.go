package services

import (
	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/models"
)

// GetProvinces is a function to get all provinces data from database
// @Summary Get all provinces
// @Description Get all provinces
// @Tags provinces
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.Province}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /api/v1/provinces [get]
func GetProvinces(c *fiber.Ctx) error {
	db := database.DBConn
	var provinces []models.Province
	db.Debug().Find(&provinces)
	// for i, s := range provinces {
	// 	fmt.Println(i, s.NameTh)
	// }
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"success":   true,
	// 	"provinces": provinces,
	// })
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all provinces.",
		Data:    provinces,
	})
}
