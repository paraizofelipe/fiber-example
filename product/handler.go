package product

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	Service Service
}

func (h ProductHandler) FindById(context *fiber.Ctx) (err error) {
	id := context.Params("id")
	product, err := h.Service.FindById(id)
	if err != nil {
		return
	}

	return context.JSON(product)
}

func (h ProductHandler) FindByTitle(context *fiber.Ctx) (err error) {
	title := context.Params("title")
	product, err := h.Service.FindByTitle(title)
	if err != nil {
		return
	}

	return context.JSON(product)
}

func (h ProductHandler) Create(context *fiber.Ctx) (err error) {
	product := Product{}
	if err = context.BodyParser(&product); err != nil {
		return
	}

	if err = h.Service.Add(product); err != nil {
		return
	}

	return context.JSON(fiber.Map{
		"status":  "success",
		"message": "created product",
	})
}

func (h ProductHandler) Update(context *fiber.Ctx) (err error) {
	id := context.Params("id")
	product := Product{}
	product.ID = id

	if err = context.BodyParser(&product); err != nil {
		return
	}

	if err = h.Service.Update(product); err != nil {
		return
	}

	return context.JSON(fiber.Map{
		"status":  "sucess",
		"message": "updated product",
	})
}

func (h ProductHandler) Delete(context *fiber.Ctx) (err error) {
	id := context.Params("id")

	if err = h.Service.DeleteById(id); err != nil {
		log.Println(err)
		return
	}

	return context.JSON(fiber.Map{
		"status":  "sucess",
		"message": "updated product",
	})
}
