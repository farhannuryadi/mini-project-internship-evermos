package response

import "time"

type UserRes struct {
	ID           uint      `json:"id"`
	Nama         string    `json:"nama"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string   	`json:"email"`
	IsAdmin      bool      `json:"is_admin"`
	IdProvinsi   string    `json:"id_provinsi"`
	IdKota       string    `json:"id_kota"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}