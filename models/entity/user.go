package entity

import "time"

type User struct {
	ID           uint					`json:"id" gorm:"primaryKey"`
	Nama         string				`json:"nama" gorm:"size:255"`
	KataSandi    string				`json:"kata_sandi" gorm:"size:255"`
	NoTelp       *string			`json:"no_telp" gorm:"size:255;unique"`
	TanggalLahir time.Time		`json:"tanggal_lahir" gorm:"type:date"`
	JenisKelamin string				`json:"jenis_kelamin" gorm:"size:255"`
	Tentang      string				`json:"tentang"`
	Pekerjaan    string				`json:"pekerjaan" gorm:"size:255,"`
	Email        *string			`json:"email" gorm:"unique;size:255"`
	IsAdmin      bool					`json:"is_admin"`
	IdProvinsi   *string			`json:"id_provinsi" gorm:"size:255"`
	IdKota       *string			`json:"id_kota" gorm:"size:255"`
	CreatedAt    time.Time		`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time		`json:"updated_at" gorm:"autoUpdateTime"`
	Trxs 				 []Trx				`json:"trxs"`
	Alamats			 []Alamat			`json:"alamats"`
	Tokos				 []Toko				`json:"tokos"`
}