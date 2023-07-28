package QuestionRouter

import (
	"go-template/middleware"

	"github.com/gofiber/fiber/v2"

	QuestionController "go-template/api/questions/controllers"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("question")
	app.Post("/create", middleware.AuthorizationRequired(), QuestionController.CreateQuestion)
}
