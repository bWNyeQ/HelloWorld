package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, f string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+f, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error: could not parse template ", err)
		return
	}

}
