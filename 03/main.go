package main

import (
	"gophercises/03/app"
	"net/http"
	"gophercises/03/app/controllers"
	"log"
)

func main() {
	s := app.New()

	s.Router.HandleFunc("/", controllers.LoadStory(s)).Methods(http.MethodGet)
	log.Fatal(s.Start())
}
