package app

import (
	"golang-microservices/mvc-pattern/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
