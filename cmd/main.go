package main

import (
	"github.com/bodecci/fintech-bank-app/cmd/api"
)

func main() {
	//migrations.Migrate()
	api.StartApi()

}
