package system_errors

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SystemErrorsHandler(ctx *fiber.Ctx, err error) error {
	var code http.ConnState
	var e = &LogicalError{}
	if errors.As(err, e) {
		code = e.code
	} else {
		return setUnhandledError(ctx, err)
	}
	if err = ctx.Status(int(code)).JSON(MapLogicalErrorToResponse(e)); err != nil {
		return setUnhandledError(ctx, errors.New("Error in setting response"))
	}
	return nil
}

func setUnhandledError(ctx *fiber.Ctx, err error) error {
	log.Printf("Unhandled error happened: %s\n", err.Error())
	return ctx.Status(http.StatusInternalServerError).JSON(MapLogicalErrorToResponse(
		NewLogicalError("Internal Server Error", http.StatusInternalServerError),
	))
}
