package services

import (
	"go-crud-api/models"
	"go-crud-api/utils"
	"go-crud-api/repositories"
)

func CreateUser(user *models.User) error {
	// 🔐 Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return repositories.CreateUser(user)
}

func GetUsers(page, limit int) ([]models.User, int64, error) {
	offset := (page - 1) * limit
	return repositories.GetUsers(offset, limit)
}

func GetUserByID(id string) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func UpdateUser(id string, updatedData *models.User) (*models.User, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = updatedData.Name
	user.Email = updatedData.Email

	if err := repositories.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id string) error {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return err
	}
	return repositories.DeleteUser(user)
}
