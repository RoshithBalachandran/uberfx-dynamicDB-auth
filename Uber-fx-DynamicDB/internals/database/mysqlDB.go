package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/roshith/dynamicDB/internals/config"
	"github.com/roshith/dynamicDB/internals/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnection(cfg *config.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.MY_USER, cfg.MY_PASS, cfg.MY_HOST, cfg.MY_PORT, cfg.MY_DB_NAME)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect MysqlDB")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("failed to automigrate User Mysql")
	}
	return db, nil
}
