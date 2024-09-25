package database

import (
	"database/sql"
	"fmt"
	"log"
	"websocket-chat/models"
)

func RegisterUser(user models.User) error {
    _, err := DB.Exec(`
        INSERT INTO users (username, email, name, password) 
        VALUES (?, ?, ?, ?)`, user.Username, user.Email, user.Name, user.Password)

    if err != nil {
        log.Printf("error when saving user: %v", err)
        return err
    }

    log.Println("user registered successfully")
    return nil
}

func FindUserBy(fieldName string, searchValue string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT name, username, email FROM users WHERE %v = ?", fieldName)

    err := DB.QueryRow(query, searchValue).Scan(&user.Name, &user.Username, &user.Email)
	if err != nil {
        if err == sql.ErrNoRows {
            return user, err
        }
        log.Printf("error when searching user by %v: %v", fieldName, err)
        return user, err
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
