package models

import (
	"fmt"

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
	fmt.Println("GetAmphures")
	db := database.DBConn
	var amphures []Amphure
	db.Unscoped().Find(&amphures)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"amphures": amphures,
	})
}
