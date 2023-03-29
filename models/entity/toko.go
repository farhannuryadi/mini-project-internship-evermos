package entity

import "time"

type Toko struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	NamaToko   string      `json:"nama_toko" form:"nama_toko" gorm:"size:255"`
	UrlFoto    string      `json:"url_foto" form:"photo" gorm:"size:255"`
	CreatedAt  time.Time   `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"-" gorm:"autoUpdateTime"`
	UserID     uint        `json:"user_id"`
	User       User        `json:"-" gorm:"foreignKey:UserID"`
	Produks    []Produk    `json:"-"`
	DetailTrxs []DetailTrx `json:"-"`
	LogProduks []LogProduk `json:"-"`
}
