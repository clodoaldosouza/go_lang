package main

import (
	"souza.com/index/alura.api.rest.gin/database"
	"souza.com/index/alura.api.rest.gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
