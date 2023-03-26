package request

type TokoReq struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   string `json:"user_id"`
}

type TokoUpdateReq struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}