package entity

type Kota struct {
	ID         string `json:"id" gorm:"size:255"`
	Name       string `json:"name" gorm:"size:255"`
	ProvinceId string `json:"province_id" gorm:"size:255"`
}
