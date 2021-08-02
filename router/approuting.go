package approuting

import (
	"cacheisking.com/handlers"
	fiber "github.com/gofiber/fiber/v2"

)


func RoutingSetup (app *fiber.App) error {
	app.Post("/setCache/temp", cachinghandler.SetCache)

	return nil
}