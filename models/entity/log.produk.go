package entity

import "time"

type LogProduk struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	NamaProduk    string      `json:"nama_produk" gorm:"size:255"`
	Slug          string      `json:"slug" gorm:"size:255"`
	HargaSeller   string      `json:"harga_seller" gorm:"size:255"`
	HargaKonsumen string      `json:"harga_konsumen" gorm:"size:255"`
	Deskripsi     string      `json:"deskripsi" gorm:"type:text"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	ProdukID      uint        `json:"produk_id"`
	TokoID        uint        `json:"toko_id"`
	CategoryID    uint        `json:"category_id"`
	DetailTrxs    []DetailTrx `json:"detail_trxs"`
}
