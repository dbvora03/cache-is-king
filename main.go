package main

import (
	"fmt"

	approuting "cacheisking.com/router"
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	approuting.RoutingSetup(app)

	fmt.Println("listening on http://0.0.0.0:1244")
	app.Listen(":1244")
}