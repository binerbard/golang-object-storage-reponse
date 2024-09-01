package main

import (
	"object_storage/internal/infrastructure/storage"
	"object_storage/pkg/routes"

	"github.com/gofiber/fiber/v2"
)


func main()  {
	storage.InitStorage()

	app := fiber.New()
	routes.SetRoutes(app)
	app.Listen(":3000")
}