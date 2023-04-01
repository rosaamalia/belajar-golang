package main

import "sesi2/router"

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)
}