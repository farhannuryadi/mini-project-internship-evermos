package entity

import "time"

type Category struct {
	ID           	uint				`json:"id" gorm:"primaryKey"`
	NamaCategory 	string			`json:"nama_category" gorm:"size:255"`
	CreatedAt    	time.Time		`json:"created_ad" gorm:"autoCreateTime"`
	UpdatedAt			time.Time		`json:"updated_at" gorm:"autoUpdateTime"`
	Produks				[]Produk		`json:"produks"`
	LogProduks		[]LogProduk	`json:"log_produks"`
}