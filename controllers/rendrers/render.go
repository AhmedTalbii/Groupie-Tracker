package rendrers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
)

// RenderPage renders a single HTML page with given data.
func RenderPage(Page string, Data any, w http.ResponseWriter) error {
	temp := template.Must(template.ParseFiles(config.Pages + Page + ".html"))
	err := temp.Execute(w, Data)
	return err
}

// MustRender checks if the given render function is valid.
func MustRender(f func(Page string, Data any, w http.ResponseWriter) error) {
	if f != nil {
		// 500 Internal server
		return
	}
}
