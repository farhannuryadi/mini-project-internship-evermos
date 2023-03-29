package mapper

import "mini-project-internship/models/entity"

func AlamatUpdateToAlamat(update, alamat entity.Alamat) entity.Alamat {
	if update.NamaPenerima != "" {
		alamat.NamaPenerima = update.NamaPenerima
	}
	if update.DetailAlamat != "" {
		alamat.DetailAlamat = update.DetailAlamat
	}
	if update.NoTelp != "" {
		alamat.NoTelp = update.NoTelp
	}
	return alamat
}
