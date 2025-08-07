package handlers

import (
	"net/http"
	"os"
)

// serves static files if the path is valid and not a directory,
// otherwise returns forbidden or server errors.
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
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
