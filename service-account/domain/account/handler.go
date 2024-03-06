package account

import (
	"fmt"
	"net/http"
	"service-account/domain/account/feature"
	"service-account/domain/account/model"
	"service-account/domain/shared/constant"
	"service-account/domain/shared/response"
	"service-account/domain/shared/validator"

	Error "service-account/domain/shared/error"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler interface {
	RegisterHandler(c *fiber.Ctx) error
	SavingHandler(c *fiber.Ctx) error
	WitdrawalHandler(c *fiber.Ctx) error
	BalanceHandler(c *fiber.Ctx) error
	HistoryHandler(c *fiber.Ctx) error
	TransferHandler(c *fiber.Ctx) error
	GenerateToken(c *fiber.Ctx) error
}

type accountHandler struct {
	accountFeature feature.AccountFeature
}

func NewAccountHandler(accountFeature feature.AccountFeature) AccountHandler {
	return &accountHandler{
		accountFeature: accountFeature,
	}
}

func (handler accountHandler) RegisterHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.CreateAccountRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.accountFeature.CreateRekeningFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler accountHandler) SavingHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.SavingAccountRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.accountFeature.SavingFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler accountHandler) WitdrawalHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.WitdrawalAccountRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.accountFeature.WitdrawalFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler accountHandler) BalanceHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	norek := c.Params("no_rekening")
	if norek == "" || norek == "0" {
		err := Error.New(ctx, constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrInvalidRequest))
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := handler.accountFeature.BalanceFeature(ctx, norek)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}

func (handler accountHandler) HistoryHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	norek := c.Params("no_rekening")
	if norek == "" || norek == "0" {
		err := Error.New(ctx, constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrInvalidRequest))
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := handler.accountFeature.HistoryFeature(ctx, norek)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}

func (handler accountHandler) TransferHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.TransferRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.accountFeature.TransferFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler accountHandler) GenerateToken(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.GenerateToken
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.accountFeature.GenerateTokenFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}
