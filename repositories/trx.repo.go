package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type TrxRepo struct {
}

func (r *TrxRepo) Create(trx entity.Trx) (entity.Trx, error) {
	err := database.DB.Debug().Create(&trx).Error
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *TrxRepo) FindById(trxId string) (entity.Trx, error) {
	var trx entity.Trx
	err := database.DB.Debug().First(&trx, "id = ?", trxId).Error
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *TrxRepo) Count(userId string) (int64, error) {
	var trx entity.Trx
	var count int64
	err := database.DB.Debug().Model(&trx).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TrxRepo) FindAllPage(userId, search string, page, limit uint64) ([]entity.Trx, error) {
	var trxs []entity.Trx
	offset := (page - 1) * limit
	err := database.DB.Raw(
		"SELECT t.* FROM trxes t, detail_trxes d, log_produks l WHERE t.id = d.trx_id AND d.log_produk_id = l.id AND l.nama_produk LIKE ? AND t.user_id = ? ORDER BY id DESC LIMIT ? OFFSET ?",
		"%"+search+"%", userId, limit, offset).Scan(&trxs).Error
	if err != nil {
		return nil, err
	}
	return trxs, err
}
