package services
import "fmt"

import (
	"errors"

	"go-crud-api/models"
	"go-crud-api/repositories"
	"go-crud-api/utils"
)

func Login(email, password string) (*models.User, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		fmt.Printf( "user1")
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(user.Password, password) {
		fmt.Printf( user.Password,password)

		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
