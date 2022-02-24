package main

import (
	"net/http"

	routes "souza.com/index/alura.web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
