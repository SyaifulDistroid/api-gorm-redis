package cmd

import (
	"context"
	"fmt"
	"service-mutasi/delivery/container"
	"service-mutasi/delivery/server"
)

func Execute() {
	// define container
	container := container.SetupContainer()

	go container.MutasiFeature.CreateMutasiFeature(context.TODO())

	// start consumer
	server := server.ServeHttp(container)
	server.Listen(fmt.Sprintf("http://localhost:%d", container.EnvironmentConfig.App.Port))
}
