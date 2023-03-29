package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
)

func TestHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AmbilDataProvinsi(c *fiber.Ctx) error {
	// Ambil data dari API
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get data from API",
		})
	}
	defer resp.Body.Close()

	var provinces []Province
	err = json.NewDecoder(resp.Body).Decode(&provinces)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse data from API",
		})
	}

	// Simpan data ke dalam database
	for _, province := range provinces {
		log.Println(province.Name)
		// err := db.Create(&province).Error
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"message": "Failed to save data to database",
		// 	})
		// }
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully saved %d provinces to database", len(provinces)),
	})
}

func TestUpload(ctx *fiber.Ctx) error {

	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		log.Println("error upload file")
	}

	var fileName string

	if file != nil {
		fileName = file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/files/%s", fileName))
		if errSaveFile != nil {
			log.Println(errSaveFile.Error())
		}
	} else {
		log.Println("tidak ada file upload")
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, "oke")
}

func TestUploadMultiple(ctx *fiber.Ctx) error {
	namaProduk := ctx.FormValue("nama_produk")

	fileNames := ctx.Locals("fileNames").([]string)
	log.Println("filenames :", fileNames)
	for _, fileName := range fileNames {
		log.Println(fileName)
	}
	log.Println(namaProduk)

	return helper.SuccessHelper(ctx, fiber.StatusOK, "oke")
}

func TestProdukFoto(c *fiber.Ctx) error {
	// Parse form data
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Get values from form data
	namaProduk := form.Value["nama_produk"]
	categoryID := form.Value["category_id"]
	hargaReseller := form.Value["harga_reseller"]
	hargaKonsumen := form.Value["harga_konsumen"]
	stok := form.Value["stok"]
	deskripsi := form.Value["deskripsi"]

	// Get files from form data
	files := form.File["photos"]

	// Store files in a slice
	var fileNames []string
	for _, file := range files {
		now := time.Now().Local().Day()
		fileName := fmt.Sprintf("%d-%v", now, file.Filename)
		// Save file to disk
		err = c.SaveFile(file, fmt.Sprintf("./public/files/%s", fileName))
		if err != nil {
			log.Println("Worning")
		}
		fileNames = append(fileNames, fileName)
	}

	// Do something with the form data and files
	// ...

	return c.JSON(fiber.Map{
		"nama produk":   namaProduk,
		"category Id":   categoryID,
		"hargaReseller": hargaReseller,
		"hargaKonsumen": hargaKonsumen,
		"stok":          stok,
		"deskripsi":     deskripsi,
		"photos":        fileNames,
	})
}
