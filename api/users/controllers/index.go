package UserController

import (
	Schema "go-template/api/users/schema"
	UserServices "go-template/api/users/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func SignUpUsers(c *fiber.Ctx) error {
	body := new(Schema.RequestSignUp)
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
	err = UserServices.SignUp(body.UserCode, body.Password, body.Firstname, body.Lastname, body.BranchCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "sign-up successfully.",
	})
}

func SignIn(c *fiber.Ctx) error {
	body := new(Schema.RequestSignIn)
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
	token, err := UserServices.SignIn(body.UserCode, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "UnAthorization",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Authorization successfully. Please save for use API later...",
		"result":  token,
	})
}
