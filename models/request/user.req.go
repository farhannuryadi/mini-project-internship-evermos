package request

type UserReq struct {
	Nama         string  `json:"nama"`
	KataSandi    string  `json:"kata_sandi"`
	NoTelp       *string `json:"no_telp"`
	TanggalLahir string  `json:"tanggal_lahir"`
	JenisKelamin string  `json:"jenis_kelamin"`
	Tentang      string  `json:"tentang"`
	Pekerjaan    string  `json:"pekerjaan"`
	Email        *string `json:"email"`
	IsAdmin      bool    `json:"is_admin"`
	IdProvinsi   *string `json:"id_provinsi"`
	IdKota       *string `json:"id_kota"`
}

type UserUpdateReq struct {
	Nama       string  `json:"nama"`
	NoTelp     *string `json:"no_telp"`
	Tentang    string  `json:"tentang"`
	Pekerjaan  string  `json:"pekerjaan"`
	Email      *string `json:"email"`
	IdProvinsi *string `json:"id_provinsi"`
	IdKota     *string `json:"id_kota"`
}