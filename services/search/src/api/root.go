package api

import (
	"github.com/A-Siam/bracker/search/src/api/routers"
	"github.com/A-Siam/bracker/search/src/common/routing"
	"github.com/gofiber/fiber/v2"
)

func InitApi() *fiber.App {
	app := fiber.New(fiber.Config{})
	routing.RegisterRouter(app, routers.UserRouter)
	return app
}
