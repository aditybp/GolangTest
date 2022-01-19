package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var SecretKey = "secret"

func MiddlewareApi(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"messages": "belum login, silahkan login",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if claims.Issuer == "1" {
		return c.Next()
	}

	return c.JSON("anda bukan admin !!")
}
