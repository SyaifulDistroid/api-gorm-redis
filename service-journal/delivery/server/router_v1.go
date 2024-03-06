package server

import (
	"service-journal/delivery/server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func routerGroupV1(app *fiber.App, handler handler) {

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,OPTION",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(middleware.Logger())
}
