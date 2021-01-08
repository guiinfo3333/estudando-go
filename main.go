package main

import (
	"estudando-go/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
