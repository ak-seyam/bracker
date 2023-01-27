package routing

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Router map[string]RouterDesc

func RegisterRouter(app *fiber.App, router Router) {
	for route, desc := range router {
		switch desc.Method {
		case Method(GET):
			registerGET(app, route, desc.Handlers)
		case Method(POST):
			registerPOST(app, route, desc.Handlers)
		case Method(PUT):
			registerPUT(app, route, desc.Handlers)
		case Method(PATCH):
			registerPATCH(app, route, desc.Handlers)
		case Method(DELETE):
			registerDELETE(app, route, desc.Handlers)
		}
	}
}

const prefix = "/api"

func registerGET(app *fiber.App, route string, handlers []fiber.Handler) {
	app.Get(prefix+route, handlers...)
}

func registerPOST(app *fiber.App, route string, handlers []fiber.Handler) {
	log.Println("register post for", prefix+route)
	app.Post(prefix+route, handlers...)
}

func registerPUT(app *fiber.App, route string, handlers []fiber.Handler) {
	app.Put(prefix+route, handlers...)
}

func registerPATCH(app *fiber.App, route string, handlers []fiber.Handler) {
	app.Patch(prefix+route, handlers...)
}

func registerDELETE(app *fiber.App, route string, handlers []fiber.Handler) {
	app.Delete(prefix+route, handlers...)
}
