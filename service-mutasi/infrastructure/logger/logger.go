package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	shared_constant "service-mutasi/domain/shared/constant"
	"service-mutasi/infrastructure/shared/constant"
	"strconv"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	useLog, XID, LogName string
)

type LogConfig struct {
	Path      string
	Prefix    string
	Extension string
}

func InitializeLogrusLogger(cfg LogConfig) {
	currentTime := time.Now()
	date := currentTime.Format("20060102")
	fileName := fmt.Sprintf("%s-%s.%s", cfg.Prefix, date, cfg.Extension)
	path := fmt.Sprintf("%s/%s", cfg.Path, fileName)
	log.SetFormatter(&log.JSONFormatter{})

	err := os.MkdirAll(filepath.Dir(path), 0770)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.SetOutput(os.Stdout)
			return
		}
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.SetOutput(os.Stdout)
	}

	LogName = path
}

func LogInfo(ctx context.Context, logtype, message string) {

	isLogErrorOnly, _ := strconv.ParseBool(os.Getenv("IS_LOG_ERROR_ONLY"))
	if !isLogErrorOnly {
		if ctx.Value(constant.SearchLogging) == nil {
			var xid string

			if ctx.Value(shared_constant.Xid) == nil {
				xid = uuid.New().String()
			} else {
				real := fmt.Sprintf("%v", ctx.Value(shared_constant.Xid).(interface{}))
				if real == "" {
					xid = uuid.New().String()
				} else {
					xid = real
				}
			}

			log.WithFields(log.Fields{
				"app_name":    os.Getenv("APP_NAME"),
				"app_version": os.Getenv("APP_VERSION"),
				"xid":         xid,
				"log_type":    logtype,
			}).Info(message)

		}
	}
}

func LogInfoRequest(ctx context.Context, logtype, message string, request interface{}) {

	isLogErrorOnly, _ := strconv.ParseBool(os.Getenv("IS_LOG_ERROR_ONLY"))
	if !isLogErrorOnly {
		if ctx.Value(constant.SearchLogging) == nil {
			var xid string

			if ctx.Value(shared_constant.Xid) == nil {
				xid = uuid.New().String()
			} else {
				real := fmt.Sprintf("%v", ctx.Value(shared_constant.Xid).(interface{}))
				if real == "" {
					xid = uuid.New().String()
				} else {
					xid = real
				}
			}

			log.WithFields(log.Fields{
				"app_name":    os.Getenv("APP_NAME"),
				"app_version": os.Getenv("APP_VERSION"),
				"xid":         xid,
				"request":     request,
				"response":    nil,
				"log_type":    logtype,
			}).Info(message)

		}
	}
}

func LogInfoResponse(ctx context.Context, logtype, message string, response interface{}) {

	isLogErrorOnly, _ := strconv.ParseBool(os.Getenv("IS_LOG_ERROR_ONLY"))
	if !isLogErrorOnly {
		if ctx.Value(constant.SearchLogging) == nil {
			var xid string

			if ctx.Value(shared_constant.Xid) == nil {
				xid = uuid.New().String()
			} else {
				real := fmt.Sprintf("%v", ctx.Value(shared_constant.Xid).(interface{}))
				if real == "" {
					xid = uuid.New().String()
				} else {
					xid = real
				}
			}

			request := ctx.Value(shared_constant.Request).(LoggerRequestData)

			log.WithFields(log.Fields{
				"app_name":    os.Getenv("APP_NAME"),
				"app_version": os.Getenv("APP_VERSION"),
				"xid":         xid,
				"request":     request,
				"response":    response,
				"log_type":    logtype,
			}).Info(shared_constant.Tdr)
		}
	}
}

func LogError(ctx context.Context, logtype, errtype, message string) {

	if ctx.Value(constant.SearchLogging) == nil {
		var xid string

		if ctx.Value(shared_constant.Xid) == nil {
			xid = uuid.New().String()
		} else {
			real := fmt.Sprintf("%v", ctx.Value(shared_constant.Xid).(interface{}))
			if real == "" {
				xid = uuid.New().String()
			} else {
				xid = real
			}
		}

		log.WithFields(log.Fields{
			"app_name":    os.Getenv("APP_NAME"),
			"app_version": os.Getenv("APP_VERSION"),
			"error_type":  errtype,
			"xid":         xid,
			"log_type":    logtype,
		}).Error(message)
	}
}
