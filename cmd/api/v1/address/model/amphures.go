package model

//Amphure is a model for amphure
type Amphure struct {
	// gorm.Model
	ID         int    `json:"id"`
	Code       string `json:"code"`
	NameTh     string `json:"name_th"`
	NameEn     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}
