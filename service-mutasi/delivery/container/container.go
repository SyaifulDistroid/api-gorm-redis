package container

import (
	"fmt"
	"log"
	"service-mutasi/config"
	mutasi_feature "service-mutasi/domain/mutasi/feature"
	"service-mutasi/domain/mutasi/model"
	mutasi_repository "service-mutasi/domain/mutasi/repository"
	database "service-mutasi/infrastructure/database"
	"service-mutasi/infrastructure/logger"
	"service-mutasi/infrastructure/redis"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	MutasiFeature     mutasi_feature.MutasiFeature
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

	db.AutoMigrate(&model.Mutasi{})

	fmt.Println("Loading redis...")
	rds, err := redis.LoadRedis(cfg.Redis)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	// initiate repository below here!
	mutasiRepository := mutasi_repository.NewMutasiRepository(db)

	fmt.Println("Loading feature's...")
	// initiate feature below here!
	mutasiFeature := mutasi_feature.NewMutasiFeature(cfg, mutasiRepository, rds)

	return Container{
		EnvironmentConfig: cfg,
		MutasiFeature:     mutasiFeature,
	}
}
