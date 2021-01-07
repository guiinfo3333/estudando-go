package routes

import (
	"guilherme/estudando-go/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)

}
