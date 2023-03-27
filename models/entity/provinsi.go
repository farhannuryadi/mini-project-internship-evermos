package entity

type Provinsi struct {
	ID   string `json:"id" gorm:"size:255"`
	Name string `json:"name" gorm:"size:255"`
}
