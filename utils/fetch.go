package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"mini-project-internship/database"
	"mini-project-internship/models/entity"

)

func FetchData() {
	provinsi := fetchProvinsi()
	fetchKota(provinsi)
}

func fetchProvinsi() []entity.Provinsi {
	urlProvinsi := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(urlProvinsi)
	if err != nil {
		log.Println("message: Failed to get data from API")
		return nil
	}
	defer resp.Body.Close()

	var provinces []entity.Provinsi
	err = json.NewDecoder(resp.Body).Decode(&provinces)
	if err != nil {
		log.Println("message: Failed to Decode")
		return nil
	}

	return provinces
}

func fetchKota(provinces []entity.Provinsi) {
	for _, provinsi := range provinces {
		database.DB.Debug().Create(&provinsi)
		urlCity := "https://www.emsifa.com/api-wilayah-indonesia/api/regencies/" + provinsi.ID + ".json"
		resp, err := http.Get(urlCity)
		log.Println(urlCity)
		if err != nil {
			log.Println("message: Failed to get data from API City")
		}
		defer resp.Body.Close()

		var citys []entity.Kota
		err = json.NewDecoder(resp.Body).Decode(&citys)
		if err != nil {
			log.Println("message: Failed to Decode City")
		}

		for _, kota := range citys {
			database.DB.Debug().Create(&kota)
		}
	}
}
