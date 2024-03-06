package server

import (
	"service-account/delivery/server/middleware"

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
	app.Post("/account/daftar", handler.accountHandler.RegisterHandler)
	app.Post("/account/generate", handler.accountHandler.GenerateToken)

	// Auth Validation
	app.Use(middleware.AuthValidations())

	v1 := app.Group("/v1")
	{
		v1.Get("/ping", handler.healthHandler.Ping)
		v1.Get("/health-check", handler.healthHandler.HealthCheck)
		tools := v1.Group("/tools")
		{
			log := tools.Group("/log")
			{
				log.Use(middleware.Logging())
				log.Post("/xid", handler.healthHandler.GetLogDataByXID)
				log.Get("/all", handler.healthHandler.GetLogData)
			}
		}

		account := v1.Group("/account")
		{
			account.Post("/tabung", handler.accountHandler.SavingHandler)
			account.Post("/tarik", handler.accountHandler.WitdrawalHandler)
			account.Post("/transfer", handler.accountHandler.TransferHandler)
			account.Get("/saldo/:no_rekening", handler.accountHandler.BalanceHandler)
			account.Get("/mutasi/:no_rekening", handler.accountHandler.HistoryHandler)
		}
	}
}
