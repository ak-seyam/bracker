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

type RouterDef struct {
	Post   *[]fiber.Handler
	Put    *[]fiber.Handler
	Get    *[]fiber.Handler
	Patch  *[]fiber.Handler
	Delete *[]fiber.Handler
}
