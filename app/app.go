package app

import (
	"golang-microservices/controllers"
	"log"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
