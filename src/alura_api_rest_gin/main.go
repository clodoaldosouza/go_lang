package main

import (
	"souza.com/index/api-go-gin/database"
	"souza.com/index/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
