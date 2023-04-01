package mapper

import "mini-project-internship/models/entity"

func ProdukToLog(prod entity.Produk) entity.LogProduk {
	var logProduk = entity.LogProduk{
		ProdukID:      prod.ID,
		TokoID:        prod.TokoID,
		CategoryID:    prod.CategoryID,
		NamaProduk:    prod.NamaProduk,
		Slug:          prod.Slug,
		HargaSeller:   prod.HargaSeller,
		HargaKonsumen: prod.HargaKonsumen,
		Deskripsi:     prod.Deskripsi,
	}

	return logProduk
}
