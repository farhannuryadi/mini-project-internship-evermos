package response

import "mini-project-internship/models/entity"

type ProdukRes struct {
	ID            uint                `json:"id"`
	NamaProduk    string              `json:"nama_produk"`
	Slug          string              `json:"slug"`
	HargaReseler  string              `json:"harga_reseler"`
	HargaKonsumen string              `json:"harga_konsumen"`
	Stok          string              `json:"stok"`
	Deskripsi     string              `json:"deskripsi"`
	Toko          entity.Toko         `json:"toko"`
	Category      entity.Category     `json:"category"`
	Photos        []entity.FotoProduk `json:"photos"`
}
