package main

import (
	"my-go-server/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
