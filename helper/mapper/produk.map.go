package mapper

import (
	"strconv"

	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
)

func ProdukReqToProduk(req request.ProdukReq) entity.Produk {
	num, _ := strconv.ParseUint(req.CategoryID, 10, 64)
	return entity.Produk{
		NamaProduk:    req.NamaProduk,
		CategoryID:    uint(num),
		HargaSeller:   req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		Slug:          req.Slug,
	}
}
