package main

import (
	"go-template/database"
	"go-template/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))
	router.SetupRoutes(app)
	app.Use(compress.New())
	app.Listen(":6000")
}
