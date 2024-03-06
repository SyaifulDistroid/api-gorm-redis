package feature

import (
	"context"
	"encoding/json"
	"fmt"
	"service-journal/config"
	"service-journal/domain/journal/model"
	repository "service-journal/domain/journal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type JournalFeature interface {
	CreateJournalFeature(ctx context.Context)
}

type journalFeature struct {
	config            config.EnvironmentConfig
	journalRepository repository.JournalRepository
	redis             *redis.Client
}

func NewJournalFeature(config config.EnvironmentConfig, journalRepository repository.JournalRepository, redis *redis.Client) JournalFeature {
	return &journalFeature{
		config:            config,
		journalRepository: journalRepository,
		redis:             redis,
	}
}

func (t *journalFeature) CreateJournalFeature(ctx context.Context) {
	for {
		// Get the JSON data from the key in Redis
		val, _ := t.redis.Get(ctx, "insert-data-key").Result()
		if val != "" {
			var data model.CreateJournalRequest
			err := json.Unmarshal([]byte(val), &data)
			if err != nil {
				fmt.Println("Error unmarshaling data:", err)
				return
			}

			err = t.journalRepository.CreateJournalRepository(ctx, &model.Journal{
				TanggalTramsaksi: time.Now(),
				NoRekeningKredit: data.NoRekeningKredit,
				NoRekeningDebit:  data.NoRekeningDebit,
				NominalKredit:    data.NominalKredit,
				NominalDebit:     data.NominalDebit,
			})

			if err != nil {
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}
