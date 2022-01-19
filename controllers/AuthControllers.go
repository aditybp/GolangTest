package controllers

import (
	"golangtest/database"
	"golangtest/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "secret"

func Welcome(c *fiber.Ctx) error {

	return c.JSON("selamat datang pada bagian admin !")
}

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	InputRaw := models.User{
		Email:    data["email"],
		Nama:     data["nama"],
		Password: Password,
	}

	database.DB.Create(&InputRaw)
	return c.JSON("data berhasil diinput !")
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"messages": "data tidak ditemukan !",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"messages": "salah password atau email !",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"messages": "tidak bisa login !",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"messages": "sukses login !",
	})
}
