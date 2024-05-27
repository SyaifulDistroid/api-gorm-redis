package middleware

import (
	"encoding/json"
	shared_constant "service-mutasi/domain/shared/constant"
	"service-mutasi/domain/shared/context"
	"service-mutasi/infrastructure/logger"
	"service-mutasi/infrastructure/shared/constant"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.CreateContext()
		ctx = context.SetFiberToContext(ctx, c)
		requestData := logger.LoggerRequestData{
			Method: c.Route().Method,
			Path:   c.OriginalURL(),
			Header: c.GetReqHeaders(),
			Body:   string(c.Request().Body()),
		}

		request, _ := json.Marshal(requestData)

		data := logger.LoggerRequestData{}
		_ = json.Unmarshal(request, &data)

		logger.LogInfoRequest(ctx, shared_constant.REQUEST, "incoming connection", data)

		ctx = context.SetRequestToContext(ctx, data)

		c.SetUserContext(ctx)
		c.Next()
		return nil
	}
}

func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.CreateContext()
		ctx = context.SetFiberToContext(ctx, c)

		ctx = context.SetCustomValueToContext(ctx, constant.SearchLogging, constant.SearchLogging)

		c.SetUserContext(ctx)
		c.Next()

		return nil
	}
}
