package models

import (
	"github.com/gofiber/fiber/v2"
	"omsoft.com/addressapi/database"
)

//Province ..
type Province struct {
	// gorm.Model
	ID          int    `json:"id"`
	Code        string `json:"code"`
	NameTh      string `json:"name_th"`
	NameEn      string `json:"name_en"`
	GeographyID int    `json:"geography_id"`
}

//GetProvinces ..
func GetProvinces(c *fiber.Ctx) error {
	db := database.DBConn
	var provinces []Province
	db.Find(&provinces)
	// for i, s := range provinces {
	// 	fmt.Println(i, s.NameTh)
	// }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"provinces": provinces,
	})
}
