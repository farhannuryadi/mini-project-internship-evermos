package entity

import "time"

type DetailTrx struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Kuantitas   int       `json:"kuantitas" gorm:"type:int"`
	HargaTotal  int       `json:"harga_total" gorm:"type:int"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TrxID       uint      `json:"trx_id"`
	LogProdukID uint      `json:"log_produk_id"`
	TokoID      uint      `json:"toko_id"`
}
