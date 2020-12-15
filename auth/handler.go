package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/paraizofelipe/fiber-example/setting"
	"github.com/paraizofelipe/fiber-example/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserService user.Service
}

// CheckPasswordHash ---
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login ---
func (h AuthHandler) Login(c *fiber.Ctx) (err error) {
	var input LoginInput

	if err = c.BodyParser(&input); err != nil {
		return
	}
	identity := input.Identity
	pass := input.Password

	userFound, err := h.UserService.FindByEmail(identity)
	if err != nil {
		return
	}

	if !CheckPasswordHash(pass, userFound.Password) {
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userFound.Username
	claims["user_id"] = userFound.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(setting.Core.Secret))
	if err != nil {
		return
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
