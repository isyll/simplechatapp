package validation

import (
	"errors"
	"regexp"
	"websocket-chat/models"
)

func ValidateUser(user models.User) error {
    if user.Name == "" {
        return errors.New("name must not be empty")
    }
    if user.Username == "" {
        return errors.New("username must not be empty")
    }
    if user.Email == "" {
        return errors.New("email must not be empty")
    }

    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !re.MatchString(user.Email) {
        return errors.New("email is invalid")
    }

    if len(user.Password) < 4 {
        return errors.New("password is too short")
    }

    return nil
}
