package container

import (
	"fmt"
	"log"
	"service-journal/config"
	journal_feature "service-journal/domain/journal/feature"
	"service-journal/domain/journal/model"
	journal_repository "service-journal/domain/journal/repository"
	database "service-journal/infrastructure/database"
	"service-journal/infrastructure/logger"
	"service-journal/infrastructure/redis"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	JournalFeature    journal_feature.JournalFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	cfg, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading logger...")
	logger.InitializeLogrusLogger(cfg.Log)

	fmt.Println("Loading database...")
	db, err := database.CreatePostgres(cfg.Database)
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&model.Journal{})

	fmt.Println("Loading redis...")
	rds, err := redis.LoadRedis(cfg.Redis)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	// initiate repository below here!
	journalRepository := journal_repository.NewJournalRepository(db)

	fmt.Println("Loading feature's...")
	// initiate feature below here!
	journalFeature := journal_feature.NewJournalFeature(cfg, journalRepository, rds)

	return Container{
		EnvironmentConfig: cfg,
		JournalFeature:    journalFeature,
	}
}
