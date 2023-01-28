package routing

import (
	"github.com/A-Siam/bracker/auth/src/common/loggers"
	"github.com/gofiber/fiber/v2"
)

type Router map[string]RouterDef

func RegisterRouter(app *fiber.App, router Router) {
	for route, def := range router {
		if def.Get != nil {
			registerGET(app, route, *def.Get)
		}
		if def.Post != nil {
			registerPOST(app, route, *def.Post)
		}
		if def.Put != nil {
			registerPUT(app, route, *def.Put)
		}
		if def.Patch != nil {
			registerPATCH(app, route, *def.Patch)
		}
		if def.Delete != nil {
			registerDELETE(app, route, *def.Delete)
		}
	}
}

const prefix = "/api"

func registerGET(app *fiber.App, route string, handlers []fiber.Handler) {
	loggers.InfoLogger.Println("register GET for", prefix+route)
	app.Get(prefix+route, handlers...)
}

func registerPOST(app *fiber.App, route string, handlers []fiber.Handler) {
	loggers.InfoLogger.Println("register POST for", prefix+route)
	app.Post(prefix+route, handlers...)
}

func registerPUT(app *fiber.App, route string, handlers []fiber.Handler) {
	loggers.InfoLogger.Println("register PUT for", prefix+route)
	app.Put(prefix+route, handlers...)
}

func registerPATCH(app *fiber.App, route string, handlers []fiber.Handler) {
	loggers.InfoLogger.Println("register PATCH for", prefix+route)
	app.Patch(prefix+route, handlers...)
}

func registerDELETE(app *fiber.App, route string, handlers []fiber.Handler) {
	loggers.InfoLogger.Println("register DELETE for", prefix+route)
	app.Delete(prefix+route, handlers...)
}
