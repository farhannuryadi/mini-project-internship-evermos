package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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