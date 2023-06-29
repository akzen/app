package main

import (
	"app/package/config"
	"app/package/handlers"
	"app/package/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/About", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
