package UserRouter

import (
	UserController "go-template/api/users/controllers"

	"go-template/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("users")
	// app.Get("/get-all", middleware.AuthorizationRequired())
	app.Post("/sign-up", middleware.AuthorizationRequired(), UserController.SignUpUsers)
	app.Post("/sign-in", UserController.SignIn)
}
