package main

import (
	"sesi2/database"
	"sesi2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}