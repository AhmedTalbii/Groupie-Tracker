package rendrers

import (
	"bytes"
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
	var buf bytes.Buffer
	
	errExec := tmp.Execute(&buf, data)
	if errExec != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	w.WriteHeader(status)
	tmp.Execute(w, data)
}
