package entity

import "time"

type Alamat struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	JudulAlamat  string    `json:"judul_alamat" gorm:"size:255"`
	NamaPenerima string    `json:"nama_penerima" gorm:"size:255" validate:"required"`
	NoTelp       string    `json:"no_telp" gorm:"size:255" validate:"required"`
	DetailAlamat string    `json:"detail_alamat" gorm:"size:255" validate:"required"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"-" gorm:"autoUpdateTime"`
	UserID       uint      `json:"-"`
	Trxs         Trx       `json:"-"`
}
