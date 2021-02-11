package models

import (
	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
)

//District ...
type District struct {
	// gorm.Model
	ID        int    `json:"id"`
	ZipCode   int    `json:"zip_code"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	AmphureID int    `json:"amphure_id"`
}

//GetDistricts ..
func GetDistricts(c *fiber.Ctx) error {
	db := database.DBConn
	var districts []District
	db.Find(&districts)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"districts": districts,
	})
}

//GetDistrictsByAmphureID ..
func GetDistrictsByAmphureID(c *fiber.Ctx) error {
	amphureID := c.Params("amphure_id", "0")
	db := database.DBConn
	var districts []District
	db.Debug().Where("amphure_id = ?", amphureID).Find(&districts)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"districts": districts,
	})
}
