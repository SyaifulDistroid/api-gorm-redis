package feature

import (
	"context"
	"encoding/json"
	"fmt"
	"service-mutasi/config"
	"service-mutasi/domain/mutasi/model"
	repository "service-mutasi/domain/mutasi/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type MutasiFeature interface {
	CreateMutasiFeature(ctx context.Context)
}

type mutasiFeature struct {
	config           config.EnvironmentConfig
	mutasiRepository repository.MutasiRepository
	redis            *redis.Client
}

func NewMutasiFeature(config config.EnvironmentConfig, mutasiRepository repository.MutasiRepository, redis *redis.Client) MutasiFeature {
	return &mutasiFeature{
		config:           config,
		mutasiRepository: mutasiRepository,
		redis:            redis,
	}
}

func (t *mutasiFeature) CreateMutasiFeature(ctx context.Context) {
	for {
		// Get the JSON data from the key in Redis
		val, _ := t.redis.Get(ctx, "insert-mutasi").Result()
		if val != "" {
			var data model.CreateMutasiRequest
			err := json.Unmarshal([]byte(val), &data)
			if err != nil {
				fmt.Println("Error unmarshaling data:", err)
				return
			}

			err = t.mutasiRepository.CreateMutasiRepository(ctx, &model.Mutasi{
				TanggalTransaksi: time.Time{},
				NoRekening:       data.NoRekening,
				JenisTransaksi:   data.JenisTransaksi,
				Nominal:          data.NoRekening,
			})

			if err != nil {
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}
