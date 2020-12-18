package main

import (
	"PROYECTintegrador/ProyectoGOI/routers"
	"os"
)

func main() {

	r := routers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run(":" + port) //"localhost:8081"
}
