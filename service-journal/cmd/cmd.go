package cmd

import (
	"context"
	"fmt"
	"service-journal/delivery/container"
	"service-journal/delivery/server"
)

func Execute() {
	// define container
	container := container.SetupContainer()

	go container.JournalFeature.CreateJournalFeature(context.TODO())

	// start consumer
	server := server.ServeHttp(container)
	server.Listen(fmt.Sprintf("http://localhost:%d", container.EnvironmentConfig.App.Port))
}
