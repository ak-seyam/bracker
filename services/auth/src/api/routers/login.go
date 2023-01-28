package routers

import (
	"github.com/A-Siam/bracker/auth/src/common/routing"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/service"
	"github.com/gofiber/fiber/v2"
)

var LoginRouter = routing.Router{
	"/login": {
		Post: &[]fiber.Handler{loginHandler},
	},
}

func loginHandler(ctx *fiber.Ctx) error {
	loginDto := &dto.UserLoginDto{}
	if err := ctx.BodyParser(loginDto); err != nil {
		return err
	}
	loginResponse, err := service.Login(*loginDto)
	if err != nil {
		return err
	}
	return ctx.JSON(loginResponse)
}
