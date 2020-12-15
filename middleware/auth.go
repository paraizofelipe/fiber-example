package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v2"
	"github.com/paraizofelipe/fiber-example/setting"
)

func Protected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte(setting.Core.Secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(context *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformad JWT" {
		return context.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}

	return context.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
