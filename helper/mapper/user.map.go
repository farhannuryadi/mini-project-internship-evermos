package mapper

import (
	// "fmt"
	"time"

	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/models/response"
)

func UserReqToUser(req request.UserReq) entity.User {
	t, _ := time.Parse("02/01/2006", req.TanggalLahir)
	t = t.Truncate(24 * time.Hour)
	return entity.User{
		Nama:         req.Nama,
		KataSandi:    req.KataSandi,
		NoTelp:       req.NoTelp,
		TanggalLahir: t,
		JenisKelamin: req.JenisKelamin,
		Tentang:      req.Tentang,
		Pekerjaan:    req.Pekerjaan,
		Email:        req.Email,
		IsAdmin:      req.IsAdmin,
		IdProvinsi:   req.IdProvinsi,
		IdKota:       req.IdKota,
	}
}

func UserToUserRes(user entity.User) response.UserRes {
	return response.UserRes{
		ID:           user.ID,
		Nama:         user.Nama,
		NoTelp:       *user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        *user.Email,
		IsAdmin:      user.IsAdmin,
		IdProvinsi:   *user.IdProvinsi,
		IdKota:       *user.IdKota,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
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
	if update.KataSandi != "" {
		user.KataSandi = update.KataSandi
	}
	return user
}

func RegisToUser(regisReq request.RegisterReq) entity.User {
	t, _ := time.Parse("02/01/2006", regisReq.TanggalLahir)
	t = t.Truncate(24 * time.Hour)
	return entity.User{
		Nama:         regisReq.Nama,
		NoTelp:       &regisReq.NoTelpon,
		Email:        &regisReq.Email,
		KataSandi:    regisReq.KataSandi,
		TanggalLahir: t,
		Pekerjaan:    regisReq.Pekerjaan,
		IdProvinsi:   &regisReq.ProvinsiId,
		IdKota:       &regisReq.KotaId,
		JenisKelamin: regisReq.JenisKelamin,
		Tentang:      regisReq.Tentang,
	}
}

func UserToLoginRes(user entity.User, provinsi entity.Provinsi, kota entity.Kota, token string) response.LoginRes {
	return response.LoginRes{
		Nama:         user.Nama,
		NoTelpon:     *user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Pekerjaan:    user.Pekerjaan,
		Email:        *user.Email,
		ProvinsiId:   provinsi,
		KotaId:       kota,
		Token:        token,
	}
}
