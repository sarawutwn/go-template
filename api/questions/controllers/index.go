package QuestionController

import (
	Schema "go-template/api/questions/schema"
	QuestionService "go-template/api/questions/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func CreateQuestion(c *fiber.Ctx) error {
	type request struct {
		Title       string               `json:"title" validate:"required"`
		BranchType  string               `json:"branch_type" validate:"required"`
		Description []Schema.Description `json:"description" validate:"required"`
	}
	body := new(request)
	c.BodyParser(&body)
	err := Validator.Struct(body)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is "+err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "validate error",
			"message": errors,
		})
	}
	question := Schema.Question{
		Title:      body.Title,
		BranchType: body.BranchType,
	}
	err = QuestionService.CreateQuestion(question, body.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "try again later...",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Create Quesetion successfully.",
	})
}

func GetQuestionByID(c *fiber.Ctx) error {
	return nil
}

func GetQuestionAll(c *fiber.Ctx) error {
	return nil
}
