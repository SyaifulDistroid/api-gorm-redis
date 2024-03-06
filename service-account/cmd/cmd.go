package cmd

import (
	"context"
	"fmt"
	"service-account/delivery/container"
	"service-account/delivery/server"
)

func Execute() {
	// define container
	container := container.SetupContainer()

	go container.AccountFeature.ShortRunningScript(context.TODO())

	// start http service
	server := server.ServeHttp(container)
	server.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
