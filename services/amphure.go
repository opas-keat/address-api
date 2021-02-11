package services

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/models"
)

// GetAmphures is a function to get all amphures data from database
// @Summary Get all amphures
// @Description Get all amphures
// @Tags amphures
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.Amphure}
// @Failure 503 {object} models.ResponseHTTP{}
// /amphures [get]
func GetAmphures(c *fiber.Ctx) error {
	db := database.DBConn
	var amphures []models.Amphure
	if res := db.Find(&amphures); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(models.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all amphures.",
		Data:    amphures,
	})
}

// GetAmphuresByProvinceID is a function to get all amphures data from database by province_id
// @Summary Get all amphures by province_id
// @Description Get all amphures by province_id
// @Tags amphures
// @Accept json
// @Produce json
// @Param province_id path int true "Province ID"
// @Success 200 {object} models.ResponseHTTP{data=[]models.Amphure}
// @Failure 503 {object} models.ResponseHTTP{}
// @Failure 404 {object} models.ResponseHTTP{}
// @Router /amphures/province/{province_id} [get]
func GetAmphuresByProvinceID(c *fiber.Ctx) error {
	provinceID := c.Params("province_id", "0")
	db := database.DBConn
	var amphures []models.Amphure

	if err := db.Where("province_id = ?", provinceID).Find(&amphures).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(models.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Amphures with Province ID %v not found.", provinceID),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(models.ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all amphures data by province_id.",
		Data:    amphures,
	})
}
