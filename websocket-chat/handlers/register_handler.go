package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"websocket-chat/database"
	"websocket-chat/models"
	"websocket-chat/utils"
	"websocket-chat/validation"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "json data is invalid"})
        return
    }

    if err := validation.ValidateUser(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    if err := uniqueConstraintsValidation(user); err != nil {
        c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
        return
    }

    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword
    if err := database.RegisterUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func uniqueConstraintsValidation(user models.User) (err error) {
    if _, err := database.FindUserByUsername(user.Username); err == nil {
        errStr := fmt.Sprintf("username '%v' already exists", user.Username)
        return errors.New(errStr)
    }

    if _, err := database.FindUserByEmail(user.Email); err == nil {
        errStr := fmt.Sprintf("email address '%v' already exists", user.Email)
        return errors.New(errStr)
    }

    return nil
}
