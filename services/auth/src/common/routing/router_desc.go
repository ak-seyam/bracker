package routing

import "github.com/gofiber/fiber/v2"

type Method int

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
)

type RouterDesc struct {
	Handlers []fiber.Handler
	Method
}
