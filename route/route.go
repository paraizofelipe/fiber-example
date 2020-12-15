package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/paraizofelipe/fiber-example/auth"
	"github.com/paraizofelipe/fiber-example/middleware"
	"github.com/paraizofelipe/fiber-example/product"
	"github.com/paraizofelipe/fiber-example/storage"
	"github.com/paraizofelipe/fiber-example/user"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	userService := user.NewUserService(storage.DB)
	productService := product.NewProductService(storage.DB)

	// Auth
	authHandler := auth.AuthHandler{
		UserService: userService,
	}
	authRoute := api.Group("/auth")
	authRoute.Post("/login", authHandler.Login)

	// Users
	userHandler := user.UserHandler{
		Service: userService,
	}
	userRoute := api.Group("/user")
	userRoute.Get("/id/:id", userHandler.FindById)
	userRoute.Get("/email/:email", userHandler.FindByEmail)
	userRoute.Post("/", userHandler.Create)
	userRoute.Delete("/:id", middleware.Protected(), userHandler.Delete)
	userRoute.Patch("/:id", middleware.Protected(), userHandler.Update)

	// Product
	productHandler := product.ProductHandler{
		Service: productService,
	}
	productRoute := api.Group("/product")
	productRoute.Get("/id/:id", productHandler.FindById)
	productRoute.Get("/title/:title", productHandler.FindByTitle)
	productRoute.Post("/", productHandler.Create)
	productRoute.Delete("/:id", middleware.Protected(), productHandler.Delete)
	productRoute.Patch("/:id", middleware.Protected(), productHandler.Update)
}
