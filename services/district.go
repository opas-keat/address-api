package services

import (
	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
	"omsoft.com/addressapi/models"
)

// GetDistricts is a function to get all districts data from database
// @Summary Get all districts
// @Description Get all districts
// @Tags districts
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.District}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /api/v1/districts [get]
func GetDistricts(c *fiber.Ctx) error {
	db := database.DBConn
	var districts []models.District
	db.Find(&districts)
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all districts.",
		Data:    districts,
	})
}

// GetDistrictsByAmphureID is a function to get all districts data from database by amphure_id
// @Summary Get all districts by amphure_id
// @Description Get all districts by amphure_id
// @Tags districts
// @Accept json
// @Produce json
// @Param amphure_id path int true "Amphure ID"
// @Success 200 {object} models.ResponseHTTP{data=[]models.District}
// @Failure 503 {object} models.ResponseHTTP{}
// @Failure 404 {object} models.ResponseHTTP{}
// @Router /api/v1/districts/amphure/{amphure_id} [get]
func GetDistrictsByAmphureID(c *fiber.Ctx) error {
	amphureID := c.Params("amphure_id", "0")
	db := database.DBConn
	var districts []models.District
	db.Debug().Where("amphure_id = ?", amphureID).Find(&districts)
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: "Success get all districts.",
		Data:    districts,
	})
}
