package models

type Theater struct {
	TheaterId     uint32 `json:"theaterId" gorm:"primary_key;"`
	TheaterName   string `json:"theaterName"`
	CityReferName string `json:"cityReferName"`
	City          City   `gorm:"foreignKey:CityReferName;References:CityName"`
}
