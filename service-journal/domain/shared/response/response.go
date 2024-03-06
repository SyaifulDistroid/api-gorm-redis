package response

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"service-journal/domain/shared/constant"
	Shared "service-journal/domain/shared/context"
	"service-journal/infrastructure/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Status       string      `json:"status"`
	Remark       string      `json:"remark,omitempty"`
	Data         interface{} `json:"data"`
	XID          interface{} `json:"xid,omitempty"`
	ResponseCode int         `json:"code"`
}

type ErrorValidation struct {
	Status string   `json:"status"`
	Remark []string `json:"remark,omitempty"`
}

type ErrorMessage struct {
	Remark string `json:"remark,omitempty"`
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {

	var (
		code = http.StatusOK
	)

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.SUCCESS,
		Remark:       msg,
		XID:          xid,
		Data:         data,
		ResponseCode: code,
	}

	logger.LogInfoResponse(c.UserContext(), constant.RESPONSE, msg, response)

	return c.Status(code).JSON(response)
}

func ResponseErrorWithContext(ctx context.Context, err error) error {

	var (
		errType    string
		statusCode = http.StatusBadRequest
	)

	logger.LogError(ctx, constant.RESPONSE, errType, err.Error())

	errSplit := strings.Split(err.Error(), ":")
	errMessage := errSplit[0]
	if strings.Contains(err.Error(), "pq:") || errType == constant.ErrDatabase {
		sqlErr := strings.Join(errSplit[1:], "")
		errMessage = strings.TrimSpace(sqlErr)
	} else if len(errSplit) > 1 {
		errMessage = errSplit[1]
	}

	errData := ErrorMessage{
		Remark: strings.TrimSpace(errMessage),
	}

	c := Shared.GetValueFiberFromContext(ctx)

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Remark:       errType,
		XID:          xid,
		Data:         errData,
		ResponseCode: statusCode,
	}

	logger.LogInfoResponse(c.UserContext(), constant.RESPONSE, err.Error(), response)

	return c.Status(statusCode).JSON(response)
}

func ResponseCustomError(c *fiber.Ctx, statusCode int, msg string, err error) error {

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Remark:       fmt.Sprintf("%s: %s", msg, err.Error()),
		Data:         nil,
		XID:          xid,
		ResponseCode: statusCode,
	}

	return c.Status(statusCode).JSON(response)
}

func ResponseValidation(c *fiber.Ctx, statusCode int, msg string, dataErr interface{}) error {
	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Remark:       msg,
		XID:          xid,
		Data:         dataErr,
		ResponseCode: statusCode,
	}

	return c.Status(statusCode).JSON(response)
}
