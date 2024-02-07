package main

import "belajar-gin-cars/routers"

func main() {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
