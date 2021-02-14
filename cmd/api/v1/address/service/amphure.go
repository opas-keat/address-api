package service

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/opas-keat/addressapi/cmd/api/v1/address/model"
	"github.com/opas-keat/addressapi/database"
)

// InitAmphure is a function to create init amphure data to database
// @Summary create init amphure
// @Description create init amphure
// @Tags amphures
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseHTTP{}
// @Failure 503 {object} model.ResponseHTTP{}
// @Router /amphures/init [post]
func InitAmphure(c *fiber.Ctx) error {
	database.InitAmphure()
	return c.JSON(model.ResponseHTTP{
		Success: true,
		Message: "Success init amphure.",
	})
}

// GetAmphures is a function to get all amphures data from database
// @Summary Get all amphures
// @Description Get all amphures
// @Tags amphures
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseHTTP{data=[]model.Amphure}
// @Failure 503 {object} model.ResponseHTTP{}
// /amphures [get]
func GetAmphures(c *fiber.Ctx) error {
	db := database.DBConn
	var amphures []model.Amphure
	if res := db.Find(&amphures); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(model.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(model.ResponseHTTP{
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
// @Success 200 {object} model.ResponseHTTP{data=[]model.Amphure}
// @Failure 503 {object} model.ResponseHTTP{}
// @Failure 404 {object} model.ResponseHTTP{}
// @Router /amphures/province/{province_id} [get]
func GetAmphuresByProvinceID(c *fiber.Ctx) error {
	provinceID := c.Params("province_id", "0")
	db := database.DBConn
	var amphures []model.Amphure

	if err := db.Where("province_id = ?", provinceID).Find(&amphures).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(model.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Amphures with Province ID %v not found.", provinceID),
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
		Message: "Success get all amphures data by province_id.",
		Data:    amphures,
	})
}
