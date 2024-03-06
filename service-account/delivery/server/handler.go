package server

import (
	"service-account/delivery/container"
	"service-account/domain/account"
	"service-account/domain/health"
)

type handler struct {
	healthHandler  health.HealthHandler
	accountHandler account.AccountHandler
}

func SetupHandler(container container.Container) handler {
	return *&handler{
		healthHandler:  health.NewHealthHandler(container.HealthFeature),
		accountHandler: account.NewAccountHandler(container.AccountFeature),
	}
}
