package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/roshith/dynamicDB/internals/config"
	"github.com/roshith/dynamicDB/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PsqlConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect MysqlDB")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("failed to automigrate User PSQL")
	}
	return db, nil
}
