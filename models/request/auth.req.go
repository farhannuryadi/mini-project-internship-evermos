package request

type LoginReq struct {
	NoTelp   string `json:"no_telp" validate:"required"`
	Password string `json:"kata_sandi" validate:"required"`
}

type RegisterReq struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NoTelpon     string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" validate:"required,email"`
	ProvinsiId   string `json:"id_provinsi"`
	KotaId       string `json:"id_kota"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
}
