package main

import (
	"project-mygram/database"
	"project-mygram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
