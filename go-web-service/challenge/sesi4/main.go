package main

import (
	"sesi4/routers"
)

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}