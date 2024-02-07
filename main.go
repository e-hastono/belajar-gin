package main

import "belajar-gin/routers"

func main() {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
