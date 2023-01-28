package routers

import (
	"errors"

	"github.com/A-Siam/bracker/search/src/common/routing"
	"github.com/A-Siam/bracker/search/src/service"
	"github.com/gofiber/fiber/v2"
)

var UserRouter = routing.Router{
	"/users": {
		Get: &[]fiber.Handler{findUserByNameHandler},
	},
}

func findUserByNameHandler(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	if name == "" {
		return errors.New("you should add name query parameter")
	}
	userDtos, err := service.FindUserByName(name)
	if err != nil {
		return err
	}
	return ctx.JSON(userDtos)
}
