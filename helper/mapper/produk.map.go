package mapper

import (
	"strconv"

	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/models/response"
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

func ProdukToProdukRes(
	produk entity.Produk, toko entity.Toko, category entity.Category,
	photos []entity.FotoProduk) response.ProdukRes {

	return response.ProdukRes{
		ID:            produk.ID,
		NamaProduk:    produk.NamaProduk,
		HargaReseler:  produk.HargaSeller,
		HargaKonsumen: produk.HargaKonsumen,
		Slug:          produk.Slug,
		Stok:          produk.Stok,
		Deskripsi:     produk.Deskripsi,
		Toko:          toko,
		Category:      category,
		Photos:        photos,
	}
}

func ProdukUpdateToProduk(update request.ProdukReq, produk entity.Produk) entity.Produk {
	categoryId, _ := strconv.ParseUint(update.CategoryID, 10, 0)
	if update.NamaProduk != "" {
		produk.NamaProduk = update.NamaProduk
	}
	if update.HargaReseller != "" {
		produk.HargaSeller = update.HargaReseller
	}
	if update.HargaKonsumen != "" {
		produk.HargaKonsumen = update.HargaKonsumen
	}
	if update.CategoryID != "" {
		produk.CategoryID = uint(categoryId)
	}
	if update.Deskripsi != "" {
		produk.Deskripsi = update.Deskripsi
	}
	if update.Slug != "" {
		produk.Slug = update.Slug
	}
	if update.Stok != "" {
		produk.Stok = update.Stok
	}
	return produk
}
