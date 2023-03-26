package entity

import "time"

type Toko struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	NamaToko   string      `json:"nama_toko" gorm:"size:255"`
	UrlFoto    string      `json:"url_foto" gorm:"size:255"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	UserID     uint        `json:"user_id"`
	User 			 User				 `gorm:"foreignKey:UserID"`
	Produks    []Produk    `json:"produks"`
	DetailTrxs []DetailTrx `json:"detail_trxs"`
	LogProduks []LogProduk `json:"log_produks"`
}