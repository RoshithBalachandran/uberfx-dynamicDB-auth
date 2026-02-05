package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/roshith/dynamicDB/internals/models"
	"github.com/roshith/dynamicDB/internals/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(name, email, password, userType string) (*models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		UserType: userType,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(email, password, userType string) (string, error) {
	user, err := s.repo.FindByEmail(email, userType)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"user_type": user.UserType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) GetProfile(id uint, userType string) (*models.User, error) {
	return s.repo.FindByID(id, userType)
}
