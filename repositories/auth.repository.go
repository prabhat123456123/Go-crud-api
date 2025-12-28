package repositories

import (
	"go-crud-api/config"
	"go-crud-api/models"
)

func FindUserByEmail(email string) (*models.User, error) {
	
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
