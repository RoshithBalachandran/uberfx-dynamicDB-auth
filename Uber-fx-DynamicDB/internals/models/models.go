package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
