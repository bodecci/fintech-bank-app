package migrations

import (
	"github.com/bodecci/fintech-bank-app/cmd/helpers"
	"github.com/bodecci/fintech-bank-app/interfaces"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := &[2]interfaces.User{
		{Username: "Bodmas", Email: "bodmas@bodmas.com"},
		{Username: "Mama", Email: "mama@mama.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{
			Type:    "Checking Account",
			Name:    string(users[i].Username + "'s" + "account"),
			Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}

	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := connectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
