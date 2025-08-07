package rendrers

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/models"
)

func ErrorPage(data models.Data, w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles(config.Pages + "error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(status)
	errExec := tmp.Execute(w, data)
	if errExec != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println(errExec)
	}
}
