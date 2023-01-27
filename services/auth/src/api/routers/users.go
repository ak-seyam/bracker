package routers

import (
	"github.com/A-Siam/bracker/auth/src/common/routing"
	"github.com/A-Siam/bracker/auth/src/dto"
	"github.com/A-Siam/bracker/auth/src/service"
	"github.com/gofiber/fiber/v2"
)

var UserRouter = routing.Router{
	"/users": {
		Handlers: []fiber.Handler{createUserHandler},
		Method:   routing.Method(routing.POST),
	},
}

func createUserHandler(ctx *fiber.Ctx) error {
	createUserDto := &dto.CreateUserDto{}
	if err := ctx.BodyParser(createUserDto); err != nil {
		return err
	}
	userDto, err := service.CreateUser(*createUserDto)
	if err != nil {
		return err
	}
	return ctx.JSON(userDto)
}
