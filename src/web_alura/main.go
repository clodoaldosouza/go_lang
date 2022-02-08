package main

import (
	"net/http"

	"souza.com/index/src/web_alura/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
