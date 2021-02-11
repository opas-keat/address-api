package services

import (
	"net/http"

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
// @Router /provinces [get]
func GetProvinces(c *fiber.Ctx) error {
	db := database.DBConn
	var provinces []models.Province
	if res := db.Find(&provinces); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(models.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	// for i, s := range provinces {
	// 	fmt.Println(i, s.NameTh)
	// }
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all provinces.",
		Data:    provinces,
	})
}

// SearchProvincesByName is a function to get all provinces data from database by thai name
// @Summary Get all provinces
// @Description Get all provinces
// @Tags provinces
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Success 200 {object} models.ResponseHTTP{data=[]models.Province}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /provinces/name/th [get]
func SearchProvincesByName(c *fiber.Ctx) error {
	nameThgo := c.Query("name", "1")
	db := database.DBConn
	var provinces []models.Province
	if res := db.Where("name_th LIKE ?", nameThgo+"%").Find(&provinces); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(models.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all provinces.",
		Data:    provinces,
	})
}
