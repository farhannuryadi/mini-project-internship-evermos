package response

import (
	"time"

	"mini-project-internship/models/entity"
)

type LoginRes struct {
	Nama         string          `json:"nama"`
	NoTelpon     string          `json:"no_telp"`
	TanggalLahir time.Time       `json:"tanggal_lahir"`
	Pekerjaan    string          `json:"pekerjaan"`
	Email        string          `json:"email"`
	ProvinsiId   entity.Provinsi `json:"id_provinsi"`
	KotaId       entity.Kota     `json:"id_kota"`
	Token        string          `json:"token"`
}
