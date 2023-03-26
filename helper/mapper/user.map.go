package mapper

import (
	// "fmt"
	"time"

	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/models/response"

)

func UserReqToUser(req request.UserReq) entity.User {
	t, _ := time.Parse("2006-01-02", req.TanggalLahir)
	// if err != nil {
	// 	fmt.Println("Invalid date format")
	// 	return nil
	// }
	t = t.Truncate(24*time.Hour)
	return entity.User{
		Nama: req.Nama,
		KataSandi: req.KataSandi,
		NoTelp: req.NoTelp,
		TanggalLahir: t,
		JenisKelamin: req.JenisKelamin,
		Tentang: req.Tentang,
		Pekerjaan: req.Pekerjaan,
		Email: req.Email,
		IsAdmin: req.IsAdmin,
		IdProvinsi: req.IdProvinsi,
		IdKota: req.IdKota,
	}
}

func UserToUserRes(user entity.User) response.UserRes {
	return response.UserRes{
	ID: user.ID,
	Nama: user.Nama, 
	NoTelp: *user.NoTelp,
	TanggalLahir: user.TanggalLahir,
	JenisKelamin: user.JenisKelamin,
	Tentang: user.Tentang,
	Pekerjaan: user.Pekerjaan,
	Email: *user.Email,
	IsAdmin: user.IsAdmin,
	IdProvinsi: *user.IdProvinsi, 
	IdKota: *user.IdKota,
	CreatedAt: user.CreatedAt,   
	UpdatedAt: user.UpdatedAt,
	}
}

func UserUpdateToUser(update request.UserUpdateReq, user entity.User) entity.User {
	if update.Nama != "" {
		user.Nama = update.Nama
	}
	if update.Tentang != "" {
		user.Tentang = update.Tentang
	}
	if update.Pekerjaan != "" {
		user.Pekerjaan = update.Pekerjaan
	}
	if update.NoTelp != nil {
		user.NoTelp = update.NoTelp
	}
	if update.Email != nil {
		user.Email = update.Email
	}
	if update.IdProvinsi != nil {
		user.IdProvinsi = update.IdProvinsi
	}
	if update.IdKota != nil {
		user.IdKota = update.IdKota
	}
	return user
}