package main

import (
	"log"
	"net/http"

	"example.com/example-web-app/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/tmp", handlers.NoTemplate)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Listening on http://localhost%s", port)
	_ = http.ListenAndServe(port, nil)
}
