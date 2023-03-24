package handlers

import (
	"fmt"
	"net/http"

	"example.com/example-web-app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "index.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := add(2, 2)
	response := fmt.Sprintf("This is the about page and 2 + 2 = %d", sum)
	fmt.Fprintf(w, response)
}

func NoTemplate(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is the home page")
}

func add(x, y int) int {
	return x + y
}
