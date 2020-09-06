package migrations

import (
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
}