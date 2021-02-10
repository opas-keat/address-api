package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
)

//Districts ..
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
	fmt.Println("GetDistricts")
	db := database.DBConn
	var districts []District
	db.Unscoped().Find(&districts)
	// for i, s := range provinces {
	// 	fmt.Println(i, s.NameTh)
	// }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"districts": districts,
	})
}
