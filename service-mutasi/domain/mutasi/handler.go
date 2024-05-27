package mutasi

import (
	"service-mutasi/domain/mutasi/feature"
)

type MutasiHandler interface {
}

type mutasiHandler struct {
	mutasiFeature feature.MutasiFeature
}

func NewMutasiHandler(mutasiFeature feature.MutasiFeature) MutasiHandler {
	return &mutasiHandler{
		mutasiFeature: mutasiFeature,
	}
}
