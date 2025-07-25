package controllers

import (
	"net/http"
	"path"

	"groupietracker/controllers/renderers"
)

func ServeStatictsController(w http.ResponseWriter, r *http.Request) {
	if path.Ext(r.URL.Path) != ".css" && path.Ext(r.URL.Path) != ".js" && path.Ext(r.URL.Path) != ".png" {
		renderers.ErrorRendrer(w, r, http.StatusForbidden, " Acces route forbidden")
		return
	}
	http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))).ServeHTTP(w, r)
}
