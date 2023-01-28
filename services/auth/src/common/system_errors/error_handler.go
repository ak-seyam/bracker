package system_errors

import (
	"errors"
	"net/http"

	"github.com/A-Siam/bracker/auth/src/common/loggers"
	"github.com/gofiber/fiber/v2"
)

func SystemErrorsHandler(ctx *fiber.Ctx, err error) error {
	var code http.ConnState
	var logicalError = &LogicalError{}
	var fiberError *fiber.Error
	if errors.As(err, logicalError) {
		code = logicalError.code
	} else if errors.As(err, &fiberError) {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return ctx.Status(fiberError.Code).SendString(err.Error())
	} else {
		return setUnhandledError(ctx, err)
	}
	if err = ctx.Status(int(code)).JSON(MapLogicalErrorToResponse(logicalError)); err != nil {
		return setUnhandledError(ctx, errors.New("Error in setting response"))
	}
	return nil
}

func setUnhandledError(ctx *fiber.Ctx, err error) error {
	loggers.ErrorLogger.Printf("Unhandled error happened: %s\n", err.Error())
	return ctx.Status(http.StatusInternalServerError).JSON(MapLogicalErrorToResponse(
		NewLogicalError("Internal Server Error", http.StatusInternalServerError),
	))
}
