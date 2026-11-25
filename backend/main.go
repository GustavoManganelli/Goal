package main

import (
	"example.com/goal/backend/db"
	"example.com/goal/backend/routes"
)

func main() {
	db.Connect()
	db.RunMigrations()
	router.Initialize()
}
