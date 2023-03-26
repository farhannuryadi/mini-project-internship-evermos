package entity

import "time"

type Alamat struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	JudulAlamat  string    `json:"judul_alamat" gorm:"size:255"`
	NamaPenerima string    `json:"nama_penerima" gorm:"size:255"`
	NoTelp       string    `json:"no_telp" gorm:"size:255"`
	DetailAlamat string    `json:"detail_alamat" gorm:"size:255"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID       uint      `json:"user_id"`
	Trxs         Trx       `json:"trxs"`
}