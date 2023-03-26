package entity

import "time"

type FotoProduk struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Url       string    `json:"url" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ProdukID  uint      `json:"produk_id"`
}