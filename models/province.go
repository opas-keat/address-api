package models

//Province is a model for province
type Province struct {
	// gorm.Model
	ID          int    `json:"id"`
	Code        string `json:"code"`
	NameTh      string `json:"name_th" example:"Bangkok"`
	NameEn      string `json:"name_en"`
	GeographyID int    `json:"geography_id"`
}
