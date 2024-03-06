package middleware

import (
	"encoding/json"
	"fmt"
	shared_constant "service-account/domain/shared/constant"
	"service-account/domain/shared/context"
	"service-account/domain/shared/response"
	"service-account/infrastructure/logger"
	"service-account/infrastructure/shared/constant"
	"strings"

	Error "service-account/domain/shared/error"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthValidations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		// Get token from header
		authToken := c.Get(constant.AUTHORIZATION)
		tokenString := strings.Replace(authToken, constant.BEARER, "", -1)

		if tokenString == "" {
			err := Error.New(ctx, constant.ErrAuth, constant.ErrAuth, fmt.Errorf(constant.ErrAuthEmpty))
			return response.ResponseErrorWithContext(ctx, err)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})
		if err != nil || !token.Valid {
			err := Error.New(ctx, constant.ErrAuth, constant.ErrAuth, fmt.Errorf(constant.ErrAuthEmpty))
			return response.ResponseErrorWithContext(ctx, err)
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			err := Error.New(ctx, constant.ErrAuth, constant.ErrAuth, fmt.Errorf(constant.ErrAuthEmpty))
			return response.ResponseErrorWithContext(ctx, err)
		}

		// Set auth token to context
		c.SetUserContext(ctx)

		c.Next()
		return nil
	}
}

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
