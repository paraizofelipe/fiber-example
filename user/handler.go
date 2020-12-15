package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service Service
}

func (h UserHandler) FindById(context *fiber.Ctx) (err error) {
	id := context.Params("id")
	user, err := h.Service.FindById(id)
	if err != nil {
		return
	}

	return context.JSON(user)
}

func (h UserHandler) FindByEmail(context *fiber.Ctx) (err error) {
	email := context.Params("email")
	user, err := h.Service.FindByEmail(email)
	if err != nil {
		return
	}

	return context.JSON(user)
}

func (h UserHandler) Create(context *fiber.Ctx) (err error) {
	user := User{}
	if err = context.BodyParser(&user); err != nil {
		return
	}

	if err = h.Service.Add(user); err != nil {
		return
	}

	return context.JSON(fiber.Map{
		"status":  "success",
		"message": "created user",
	})
}

func (h UserHandler) Update(context *fiber.Ctx) (err error) {
	id := context.Params("id")
	user := User{}
	user.ID = id

	if err = context.BodyParser(&user); err != nil {
		return
	}

	if err = h.Service.Update(user); err != nil {
		return
	}

	return context.JSON(fiber.Map{
		"status":  "sucess",
		"message": "updated user",
	})
}

func (h UserHandler) Delete(context *fiber.Ctx) (err error) {
	id := context.Params("id")

	if err = h.Service.DeleteById(id); err != nil {
		log.Println(err)
		return
	}

	return context.JSON(fiber.Map{
		"status":  "sucess",
		"message": "updated user",
	})
}
