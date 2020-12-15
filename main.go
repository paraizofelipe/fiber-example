package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/paraizofelipe/fiber-example/errors"
	"github.com/paraizofelipe/fiber-example/route"
	"github.com/paraizofelipe/fiber-example/setting"
	"github.com/paraizofelipe/fiber-example/storage"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler:          errors.ErrorHandler,
		CaseSensitive:         true,
		StrictRouting:         true,
		ServerHeader:          "Fiber",
		DisableStartupMessage: false,
	})
	app.Use(cors.New())

	storage.ConnectDB()

	route.SetupRoutes(app)
	log.Fatal(app.Listen(setting.Core.APIURL))

	defer storage.DB.Close()
}
