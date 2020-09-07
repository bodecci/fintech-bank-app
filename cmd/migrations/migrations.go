package migrations

import (
	"github.com/bodecci/fintech-bank-app/cmd/helpers"
	"github.com/bodecci/fintech-app/cmd/helpers"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email 	 string
	Password string
}

type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "Bodmas", Email: "bodmas@bodmas.com"},
		{Username: "Mama", Email: "mama@mama.com"},
	}

	for i:= 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email:user[i].Email, Password: generatedPassword}
		db.Create(&user)
	}

	defer db.Close
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&Users{}, &Account{})
	defer db.Close()

	createAccounts()
}