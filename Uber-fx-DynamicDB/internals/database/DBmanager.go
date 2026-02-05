package database

import (
	"errors"
	"log"

	"github.com/roshith/dynamicDB/internals/config"
	"gorm.io/gorm"
)

type DBManager struct {
	UserDB  *gorm.DB
	StaffDB *gorm.DB
}

// Database Manager
func DatabaseManager(cfg *config.Config) (*DBManager, error) {
	user, err := PsqlConnection(cfg)
	if err != nil {
		return nil, errors.New("failed to connect psql db user")
	}

	staff, err := MysqlConnection(cfg)

	if err != nil {
		return nil, errors.New("failed to connect mysql db staff")
	}

	return &DBManager{
		StaffDB: staff,
		UserDB:  user,
	}, nil
}

func (m *DBManager) GetDB(usertype string) *gorm.DB {
	if usertype == "staff" {
		log.Println("mysql db connected for staff")
		return m.StaffDB
	}
	log.Println("PSQL db connected for staff")

	return m.UserDB

}
