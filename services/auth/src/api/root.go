package api

import (
	"github.com/A-Siam/bracker/auth/src/api/routers"
	"github.com/A-Siam/bracker/auth/src/common/routing"
	"github.com/A-Siam/bracker/auth/src/common/system_errors"
	"github.com/gofiber/fiber/v2"
)

func InitApi() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: system_errors.SystemErrorsHandler,
	})
	routing.RegisterRouter(app, routers.UserRouter)
	routing.RegisterRouter(app, routers.LoginRouter)
	return app
}
