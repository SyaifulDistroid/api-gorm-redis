package feature

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"service-account/config"
	"service-account/domain/health/constant"
	"service-account/domain/health/helper"
	"service-account/domain/health/model"
	"service-account/domain/health/repository"
	shared_constant "service-account/domain/shared/constant"
	Error "service-account/domain/shared/error"
	"service-account/infrastructure/logger"
	"strings"

	"github.com/google/uuid"
)

type HealthFeature interface {
	HealthCheck(ctx context.Context) (resp model.HealthCheck, err error)
	GetLogByXID(ctx context.Context, request *model.LogRequest) (resp *model.LogDetailResponse, err error)
	GetLogAll(ctx context.Context) (resp *model.LogDetailResponse, err error)
}

type healthFeature struct {
	config           config.EnvironmentConfig
	healthRepository repository.HealthRepository
}

func NewHealthFeature(config config.EnvironmentConfig, healthRepository repository.HealthRepository) HealthFeature {
	return &healthFeature{
		config:           config,
		healthRepository: healthRepository,
	}
}

func (hf healthFeature) HealthCheck(ctx context.Context) (resp model.HealthCheck, err error) {

	var (
		status   = constant.HEALTHY
		dbstatus = constant.CONNECTED
	)

	resp = model.HealthCheck{
		AppDetail: model.AppDetail{
			Name:    hf.config.App.Name,
			Version: hf.config.App.Version,
		},
		DatabaseDetail: model.DatabaseDetail{
			Dialect: hf.config.Database.Dialect,
			Status:  dbstatus,
		},
		Status: status,
	}

	return

}

func (hf healthFeature) GetLogByXID(ctx context.Context, request *model.LogRequest) (resp *model.LogDetailResponse, err error) {

	var (
		logDetail model.LogDetailResponse
	)

	uuid, err := uuid.Parse(request.Xid)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrInvalidXID, errors.New(fmt.Sprintf(constant.ErrInvalidXIDWithError, err.Error())))
		return
	}

	logDetail.Xid = &uuid

	logs, err := helper.ReadLines(logger.LogName)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrLogDataNotFound, errors.New(fmt.Sprintf(constant.ErrLogDataNotFoundWithError, logger.LogName, err.Error())))
		return
	}

	for _, log := range logs {
		if strings.Contains(log, uuid.String()) {
			logDetail.Contents = append(logDetail.Contents, log)
		}
	}

	resp = &logDetail

	return
}

func (hf healthFeature) GetLogAll(ctx context.Context) (resp *model.LogDetailResponse, err error) {

	var (
		logDetail model.LogDetailResponse
	)

	content, err := ioutil.ReadFile(logger.LogName)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrLogDataNotFound, errors.New(fmt.Sprintf(constant.ErrLogDataNotFoundWithError, logger.LogName, err.Error())))
		return
	}

	logDetail.Contents = append(logDetail.Contents, string(content))

	resp = &logDetail

	return
}
