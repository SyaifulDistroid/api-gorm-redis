package middleware

import (
	"encoding/json"
	shared_constant "service-journal/domain/shared/constant"
	"service-journal/domain/shared/context"
	"service-journal/infrastructure/logger"
	"service-journal/infrastructure/shared/constant"

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
