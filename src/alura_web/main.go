package main

import (
	"net/http"

	"souza.com/index/src/alura_web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
