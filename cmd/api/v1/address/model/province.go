package model

//Province is a model for province
type Province struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	NameTh      string `json:"name_th"`
	NameEn      string `json:"name_en"`
	GeographyID int    `json:"geography_id"`
}
