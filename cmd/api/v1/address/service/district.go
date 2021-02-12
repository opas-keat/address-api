package service

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/opas-keat/addressapi/cmd/api/v1/address/model"
	"github.com/opas-keat/addressapi/database"
)

// GetDistricts is a function to get all districts data from database
// @Summary Get all districts
// @Description Get all districts
// @Tags districts
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseHTTP{data=[]model.District}
// @Failure 503 {object} model.ResponseHTTP{}
// //districts [get]
func GetDistricts(c *fiber.Ctx) error {
	db := database.DBConn
	var districts []model.District
	if res := db.Find(&districts); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(model.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(model.ResponseHTTP{
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
// @Success 200 {object} model.ResponseHTTP{data=[]model.District}
// @Failure 503 {object} model.ResponseHTTP{}
// @Failure 404 {object} model.ResponseHTTP{}
// @Router /districts/amphure/{amphure_id} [get]
func GetDistrictsByAmphureID(c *fiber.Ctx) error {
	amphureID := c.Params("amphure_id", "0")
	db := database.DBConn
	var districts []model.District
	if err := db.Where("amphure_id = ?", amphureID).Find(&districts).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(model.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Districts with Amphure ID %v not found.", amphureID),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(model.ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}
	return c.JSON(model.ResponseHTTP{
		Success: true,
		Message: "Success get all districts.",
		Data:    districts,
	})
}
