package api

import "github.com/gofiber/fiber/v2"

func GetApiApp() *fiber.App {
	app := fiber.New()

	app.Post("/", ping)
	app.Post("/api", api)
	app.Post("/rpc", rpc)
	app.Post("/kafka", kafka)
	return app
}
