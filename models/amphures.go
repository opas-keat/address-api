package models

import (
	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
)

//Amphure ..
type Amphure struct {
	// gorm.Model
	ID         int    `json:"id"`
	Code       string `json:"code"`
	NameTh     string `json:"name_th"`
	NameEn     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}

//GetAmphures ..
func GetAmphures(c *fiber.Ctx) error {
	db := database.DBConn
	var amphures []Amphure
	db.Find(&amphures)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"amphures": amphures,
	})
}

//GetAmphuresByProvinceID ..
func GetAmphuresByProvinceID(c *fiber.Ctx) error {
	provinceID := c.Params("province_id", "0")
	db := database.DBConn
	var amphures []Amphure
	db.Debug().Where("province_id = ?", provinceID).Find(&amphures)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"amphures": amphures,
	})
}
