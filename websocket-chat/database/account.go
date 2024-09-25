package database

import (
	"fmt"
	"websocket-chat/initializers"
	"websocket-chat/models"
)

func RegisterUser(user models.User) error {
    result := initializers.DB.Create(&user)
    return result.Error
}

func FindUserBy(fieldName string, searchValue string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("%v = ?", fieldName)

    result := initializers.DB.Where(query, searchValue).First(&user)

    if result.Error != nil {
        return user, result.Error
    }

    return user, nil
}

func FindUserByUsername(username string) (models.User, error) {
    user, err := FindUserBy("username", username)
	return user, err
}

func FindUserByEmail(email string) (models.User, error) {
    user, err := FindUserBy("email", email)
	return user, err
}
