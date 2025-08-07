package rendrers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
)

// RenderPage renders a single HTML page with given data.
func RenderPage(Page string, Data any, w http.ResponseWriter) error {
	Templ := template.Must(template.ParseFiles(config.Pages + Page + ".html"))
	err := Templ.Execute(w, Data)
	return err
}

// MustRender checks if the given render function is valid.
func MustRender(Page string, Data any, w http.ResponseWriter) {
	RenderPage(Page, Data, w)
}
