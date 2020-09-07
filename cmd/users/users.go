package users

import (
	"github.com/bodecci/fintech-bank-app/cmd/helpers"
	interfaces "github.com/bodecci/fintech-bank-app/cmd/interfaces"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {
	// create db connection. look for user with username from the function params
	db := helpers.ConnectDB()
	user := &interfaces.User{}

	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	// we verify the user by using the bcrypt.CompareHashPassword
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr != nil && passErr == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"message": "Wrong password"}
	}

}
