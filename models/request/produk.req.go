package request

type ProdukReq struct {
	NamaProduk    string `json:"nama_produk" form:"nama_produk"`
	CategoryID    string `json:"category_id" form:"category_id"`
	HargaReseller string `json:"harga_reseller" form:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen" form:"harga_konsumen"`
	Stok          string `json:"stok" form:"stok"`
	Slug          string `json:"slug" form:"slug"`
	Deskripsi     string `json:"deskripsi" form:"deskripsi"`
}
