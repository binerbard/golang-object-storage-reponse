package routes

import (
	"object_storage/pkg/handle"

	"github.com/gofiber/fiber/v2"
)

var RouteList map[string]string = map[string]string{
	"store": "/store",
	"show": "/show/:filename",
}


func SetRoutes(r fiber.Router) {
	router := r.Group("/")
	router.Post(RouteList["store"], handle.Store)
	router.Get(RouteList["show"], handle.Show)
}