package server

import (
	"service-mutasi/delivery/container"
	"service-mutasi/domain/mutasi"
)

type handler struct {
	mutasiHandler mutasi.MutasiHandler
}

func SetupHandler(container container.Container) handler {
	return *&handler{
		mutasiHandler: mutasi.NewMutasiHandler(container.MutasiFeature),
	}
}
