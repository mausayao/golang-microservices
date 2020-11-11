package routes

import (
	"golang-microservices/controllers"
	"net/http"
)

func RoutesLoad() {
	http.HandleFunc("/users", controllers.GetUser)
}
