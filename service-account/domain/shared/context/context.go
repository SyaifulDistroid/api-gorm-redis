package context

import (
	"context"
	"service-account/domain/shared/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateContext() context.Context {
	ctx := context.Background()
	uuid := uuid.New()
	return context.WithValue(ctx, constant.Xid, uuid)
}

func SetRequestToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, constant.Request, val)
}

func SetCustomValueToContext(ctx context.Context, key string, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

func SetFiberToContext(ctx context.Context, c *fiber.Ctx) context.Context {
	return context.WithValue(ctx, constant.FiberContext, c)
}

func GetValueFiberFromContext(ctx context.Context) *fiber.Ctx {
	return ctx.Value(constant.FiberContext).(*fiber.Ctx)
}
