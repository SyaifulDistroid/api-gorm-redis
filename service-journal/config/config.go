package config

import (
	"fmt"
	"os"
	database "service-journal/infrastructure/database"
	"service-journal/infrastructure/logger"
	"service-journal/infrastructure/redis"
	shared_constant "service-journal/infrastructure/shared/constant"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
	Redis    redis.RedisConfig
	Log      logger.LogConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

type Log struct {
	Path      string
	Prefix    string
	Extension string
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf(shared_constant.ErrConvertStringToInt, err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf(shared_constant.ErrConvertStringToInt, err)
		return
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.DatabaseConfig{
			Dialect:        os.Getenv("POSTGRES_DIALECT"),
			Host:           os.Getenv("POSTGRES_HOST"),
			Name:           os.Getenv("POSTGRES_NAME"),
			Username:       os.Getenv("POSTGRES_USERNAME"),
			Password:       os.Getenv("POSTGRES_PASSWORD"),
			Port:           os.Getenv("POSTGRES_PORT"),
			SetMaxIdleConn: os.Getenv("DB_SET_MAX_IDLE_CONN"),
			SetMaxOpenConn: os.Getenv("DB_SET_MAX_OPEN_CONN"),
			SetMaxIdleTime: os.Getenv("DB_SET_MAX_IDLE_TIME"),
			SetMaxLifeTime: os.Getenv("DB_SET_MAX_LIFE_TIME"),
		},
		Redis: redis.RedisConfig{
			Adress:   os.Getenv("REDIS_ADDRES"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Database: 0,
		},
		Log: logger.LogConfig{
			Path:      os.Getenv("LOG_PATH"),
			Prefix:    os.Getenv("LOG_PREFIX"),
			Extension: os.Getenv("LOG_EXT"),
		},
	}

	return
}
