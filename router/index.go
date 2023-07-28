package router

import (
	QuestionRouter "go-template/api/questions"
	UserRouter "go-template/api/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())
	UserRouter.SetupRoutes(api)
	QuestionRouter.SetupRoutes(api)
}
