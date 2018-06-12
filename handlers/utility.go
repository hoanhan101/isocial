package handlers

import (
	"html/template"
	"log"
	"net/http"
	ttemplate "text/template"
)

// RenderTemplate is responsible for rendering data object into template file.
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

// RenderUnsafeTemplate is responsible for rendering data object into template file
// in an unsafe manner.
func RenderUnsafeTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := ttemplate.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	w.Header().Set("X-XSS-Protection", "0")
	t.Execute(w, templateData)
}
