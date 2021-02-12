package service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/opas-keat/addressapi/cmd/api/v1/address/model"
	"github.com/opas-keat/addressapi/database"
)

// GetProvinces is a function to get all provinces data from database
// @Summary Get all provinces
// @Description Get all provinces
// @Tags provinces
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseHTTP{data=[]model.Province}
// @Failure 503 {object} model.ResponseHTTP{}
// @Router /provinces [get]
func GetProvinces(c *fiber.Ctx) error {
	db := database.DBConn
	var provinces []model.Province
	if res := db.Find(&provinces); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(model.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	// for i, s := range provinces {
	// 	fmt.Println(i, s.NameTh)
	// }
	return c.JSON(model.ResponseHTTP{
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
// @Success 200 {object} model.ResponseHTTP{data=[]model.Province}
// @Failure 503 {object} model.ResponseHTTP{}
// @Router /provinces/name/th [get]
func SearchProvincesByName(c *fiber.Ctx) error {
	nameThgo := c.Query("name", "1")
	db := database.DBConn
	var provinces []model.Province
	if res := db.Where("name_th LIKE ?", nameThgo+"%").Find(&provinces); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(model.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(model.ResponseHTTP{
		Success: true,
		Message: "Success get all provinces.",
		Data:    provinces,
	})
}
