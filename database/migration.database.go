package database

import (
	"fmt"
	"log"

	"mini-project-internship/models/entity"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&entity.User{},
		&entity.Alamat{},
		&entity.Toko{},
		&entity.Category{},
		&entity.Produk{},
		&entity.FotoProduk{},
		&entity.Trx{},
		&entity.DetailTrx{},
		&entity.LogProduk{},
		&entity.Provinsi{},
		&entity.Kota{},
	)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("success database migration")
}
