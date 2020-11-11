package app

import (
	"golang-microservices/routes"
	"log"
	"net/http"
)

func StartApp() {
	routes.RoutesLoad()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
