package main

import (
	"github.com/ypgyan/api-go-gin/database"
	"github.com/ypgyan/api-go-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
