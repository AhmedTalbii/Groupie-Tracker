package renderers

import (
	"log"
	"net/http"
	"text/template"

	"groupietracker/config"
)

func PageRender(w http.ResponseWriter, r *http.Request, path string, data any) {
	tmpl := template.Must(template.New("Layout").ParseFiles(
		config.Layout+"layout.html",
		config.Pages+path,
	))

	tmpl = template.Must(tmpl.ParseGlob(config.Sections + "*.html"))
	tmpl = template.Must(tmpl.ParseGlob(config.Components + "*.html"))
	errExecuteTemplate := tmpl.ExecuteTemplate(w, "Layout", data)
	if errExecuteTemplate != nil {
		log.Fatal("Error executing the template :", errExecuteTemplate)
	}
}
