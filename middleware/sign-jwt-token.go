package middleware

import (
	"go-template/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = config.GetEnvConfig("SECRET_KEY")

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: successHandler,
		ErrorHandler:   errorHandler,
		SigningKey:     []byte(secretKey),
		SigningMethod:  "HS256",
	})
}

func successHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["question_user_id"].(string)
	c.Locals("question_user_id", ID)
	return c.Next()
}
func errorHandler(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}
