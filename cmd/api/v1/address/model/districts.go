package model

//District is a model for district
type District struct {
	ID        int    `json:"id"`
	ZipCode   int    `json:"zip_code"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	AmphureID int    `json:"amphure_id"`
}
