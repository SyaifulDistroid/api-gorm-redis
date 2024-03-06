package container

import (
	"fmt"
	"log"
	"service-account/config"
	account_feature "service-account/domain/account/feature"
	"service-account/domain/account/model"
	account_repository "service-account/domain/account/repository"
	health_feature "service-account/domain/health/feature"
	health_repository "service-account/domain/health/repository"
	database "service-account/infrastructure/database"
	"service-account/infrastructure/logger"
	"service-account/infrastructure/redis"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	HealthFeature     health_feature.HealthFeature
	AccountFeature    account_feature.AccountFeature
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

	db.AutoMigrate(&model.Account{}, &model.History{})

	fmt.Println("Loading redis...")
	rds, err := redis.LoadRedis(cfg.Redis)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	// initiate repository below here!
	healthRepository := health_repository.NewHealthRepository(db)
	accountRepository := account_repository.NewAccountRepository(db)

	fmt.Println("Loading feature's...")
	// initiate feature below here!
	healthFeature := health_feature.NewHealthFeature(cfg, healthRepository)
	accountFeature := account_feature.NewAccountFeature(cfg, accountRepository, rds)

	return Container{
		EnvironmentConfig: cfg,
		HealthFeature:     healthFeature,
		AccountFeature:    accountFeature,
	}
}
