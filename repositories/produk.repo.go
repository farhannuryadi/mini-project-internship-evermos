package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type ProdukRepo struct {
}

func (r *ProdukRepo) Create(produk entity.Produk) (entity.Produk, error) {
	err := database.DB.Debug().Create(&produk).Error
	if err != nil {
		return produk, err
	}
	return produk, nil
}

func (r *ProdukRepo) FindById(produkId string) (entity.Produk, error) {
	var produk entity.Produk
	err := database.DB.Debug().First(&produk, produkId).Error
	if err != nil {
		return produk, err
	}
	return produk, nil
}

func (r *ProdukRepo) Count() (int64, error) {
	var produk entity.Produk
	var count int64
	err := database.DB.Model(&produk).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProdukRepo) FindAllPage(
	namaProduk string, categoryID string, tokoID string, maxHarga string,
	minHarga string, limit uint, page uint) ([]entity.Produk, error) {

	var produks []entity.Produk
	offset := (page - 1) * limit
	query := "SELECT * FROM produks WHERE nama_produk LIKE ?"
	values := []interface{}{"%" + namaProduk + "%"}

	if categoryID != "" {
		query += " AND category_id = ?"
		values = append(values, categoryID)
	}

	if tokoID != "" {
		query += " AND toko_id = ?"
		values = append(values, tokoID)
	}

	if maxHarga != "" {
		query += " AND harga_konsumen <= ?"
		values = append(values, maxHarga)
	}

	if minHarga != "" {
		query += " AND harga_konsumen >= ?"
		values = append(values, minHarga)
	}

	query += " LIMIT ? OFFSET ?"
	values = append(values, limit, offset)

	result := database.DB.Debug().Raw(query, values...).Scan(&produks)

	if result.Error != nil {
		return nil, result.Error
	}

	return produks, nil
}

func (r *ProdukRepo) Update(produk entity.Produk) (entity.Produk, error) {
	err := database.DB.Debug().Save(&produk).Error
	if err != nil {
		return produk, err
	}
	return produk, nil
}

func (r *ProdukRepo) Delete(produk entity.Produk) (bool, error) {
	err := database.DB.Debug().Delete(&produk).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
