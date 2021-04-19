package main

import (
	"golang-transaction/model"
	"golang-transaction/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)
}
