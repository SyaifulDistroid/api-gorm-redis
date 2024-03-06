package error

import (
	"context"
	"errors"
	"fmt"
	"service-journal/domain/shared/constant"
	"service-journal/infrastructure/logger"
	"strings"
)

func New(ctx context.Context, tipe string, msg string, err error) error {
	logger.LogError(ctx, constant.Err, tipe, fmt.Sprintf(msg+": "+err.Error()))
	return fmt.Errorf("%s | %s: %w", tipe, msg, err)
}

func TrimMesssage(err error) (tipe string, newErr error) {
	errs := strings.Split(err.Error(), "|")
	tipe = strings.TrimSpace(errs[0])

	newErr = errors.New(strings.TrimSpace(errs[1]))
	if len(errs)-1 == 2 {
		newErr = errors.New(strings.TrimSpace(errs[2]))
	} else if len(errs) > 1 {
		newErr = errors.New(strings.TrimSpace(errs[1]))
	}

	return
}

func LoopErrorFormat(i int, errStr string) string {
	err := fmt.Sprintf("row %d err: %s", i+1, errStr)
	return err
}
