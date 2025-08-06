package handlers

import (
	"net/http"
	"os"
)

// Serve the statics css files + every asset
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	// if directory acces forbidden
	info, err := os.Stat(r.URL.Path)
	if err != nil {
		// 500 intrnal server
		return
	}
	if info.IsDir() {
		// 403 Forbidden
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
