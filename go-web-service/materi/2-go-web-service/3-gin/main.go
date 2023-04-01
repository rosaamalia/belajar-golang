package main

import "3-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}