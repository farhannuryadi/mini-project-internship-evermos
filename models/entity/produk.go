package entity

import "time"

type Produk struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	NamaProduk    string       `json:"nama_produk" gorm:"size:255"`
	Slug          string       `json:"slug" gorm:"size:255"`
	HargaSeller   string       `json:"harga_seller" gorm:"size:255"`
	HargaKonsumen string       `json:"harga_konsumen" gorm:"size:255"`
	Stok          string       `json:"stok" gorm:"type:int"`
	Deskripsi     string       `json:"deskripsi" gorm:"type:text"`
	CreatedAt     time.Time    `json:"-" gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `json:"-" gorm:"autoUpdateTime"`
	TokoID        uint         `json:"toko_id"`
	CategoryID    uint         `json:"category_id"`
	FotoProduks   []FotoProduk `json:"-"`
	LogProduks    []LogProduk  `json:"-"`
}
