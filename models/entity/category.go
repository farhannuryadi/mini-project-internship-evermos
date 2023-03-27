package entity

import "time"

type Category struct {
	ID           	uint				`json:"id" gorm:"primaryKey"`
	NamaCategory 	string			`json:"nama_category" gorm:"size:255"`
	CreatedAt    	time.Time		`json:"-" gorm:"autoCreateTime"`
	UpdatedAt			time.Time		`json:"-" gorm:"autoUpdateTime"`
	Produks				[]Produk		`json:"-"`
	LogProduks		[]LogProduk	`json:"-"`
}