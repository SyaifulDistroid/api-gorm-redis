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
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	// create database
	dbConn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
