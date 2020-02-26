package app

import (
	"net/http"

	"github.com/rennanbadaro/go-mvc-webserver/controllers"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
