package request

type DetailTrxReq struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}

type TrxReq struct {
	MethodBayar string         `json:"method_bayar"`
	AlamatKirim int            `json:"alamat_kirim"`
	DetailTrx   []DetailTrxReq `json:"detail_trx"`
}
