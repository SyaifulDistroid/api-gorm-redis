package pgpool

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Name     string
	Username string
	Password string
	Port     string

	SetMaxIdleConn string
	SetMaxOpenConn string
	SetMaxIdleTime string
	SetMaxLifeTime string
}

func CreatePostgres(config DatabaseConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Dialect,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	dbConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
