package entity

import "time"

type Trx struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	HargaTotal  uint        `json:"harga_total"`
	KodeInvoice *string     `json:"kode_invoice" gorm:"size:255"`
	MethodBayar *string     `json:"method_Bayar" gorm:"size:255"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	UserID      uint        `json:"user_id"`
	AlamatID    uint        `json:"alamat_id"`
	DetailTrxs  []DetailTrx `json:"detail_trxs"`
}
