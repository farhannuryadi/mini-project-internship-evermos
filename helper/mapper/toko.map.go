package mapper

import (
	"strconv"

	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
)

func TokoReqToToko(req request.TokoReq) entity.Toko {
	num, _ := strconv.Atoi(req.UserID)
	userId := uint(num)
	return entity.Toko{
		NamaToko: req.NamaToko,
		UrlFoto:  req.UrlFoto,
		UserID:   userId,
	}
}

func TokoUpdateToToko(update entity.Toko, toko entity.Toko) entity.Toko {
	if update.NamaToko != "" {
		toko.NamaToko = update.NamaToko
	}
	if update.UrlFoto != "" {
		toko.UrlFoto = update.UrlFoto
	}
	return toko
}
