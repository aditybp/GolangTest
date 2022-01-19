package controllers

import (
	"fmt"
	"golangtest/database"
	"golangtest/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduk(c *fiber.Ctx) error {

	var produk []models.Produk
	database.DB.Find(&produk)
	if len(produk) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"data":   nil,
		})
	}
	return c.JSON(&produk)
}

func AddProduk(c *fiber.Ctx) error {
	produk := new(models.Produk)

	if err := c.BodyParser(produk); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	InputRaw := models.Produk{
		NamaBarang:   "NamaBarang",
		JumlahBarang: "JumlahBarang",
		KodeBarang:   "KodeBarang",
	}

	database.DB.Create(&InputRaw)
	return c.JSON(&InputRaw)
}

func UpdateProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	produk := new(models.Produk)
	database.DB.First(&produk, id)

	if produk.NamaBarang == "" {
		return c.Status(500).SendString("produk tidak tersedia")
	}

	if err := c.BodyParser(&produk); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&produk)
	return c.JSON(&produk)
}

func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk
	database.DB.First(&produk, id)

	if produk.NamaBarang == "" {
		c.Status(500).SendString("tidak ada barang")
	}

	database.DB.Delete(&produk)
	return c.JSON("data telah terhapus")
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk
	database.DB.Find(&produk, id)
	return c.JSON(&produk)
}

func UploadPhoto(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.JSON(fiber.Map{
			"messages": "upload error",
		})
	}

	return c.SaveFile(file, fmt.Sprintf("./images/%s", file.Filename))
}
