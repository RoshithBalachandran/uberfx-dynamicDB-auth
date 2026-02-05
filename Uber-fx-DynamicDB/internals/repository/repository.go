package repository

import (
	"github.com/roshith/dynamicDB/internals/database"
	"github.com/roshith/dynamicDB/internals/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email, userType string) (*models.User, error)
	FindByID(id uint, userType string) (*models.User, error)
}

type userRepo struct {
	dbManager *database.DBManager
}

func NewRepository(dbManager *database.DBManager) UserRepository {
	return &userRepo{dbManager: dbManager}
}

func (r *userRepo) Create(user *models.User) error {
	db := r.dbManager.GetDB(user.UserType)
	return db.Create(user).Error
}

func (r *userRepo) FindByEmail(email, userType string) (*models.User, error) {
	db := r.dbManager.GetDB(userType)
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByID(id uint, userType string) (*models.User, error) {
	db := r.dbManager.GetDB(userType)

	var user models.User
	err := db.First(&user, id).Error
	return &user, err
}
