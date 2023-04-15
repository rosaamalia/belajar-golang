package main

import (
	"sesi4/router"
	"sesi4/database"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}