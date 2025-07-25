package renderers

import (
	"net/http"
	"strconv"
)

func ErrorRendrer(w http.ResponseWriter, r *http.Request, status int, err string) {
	PageRender(w, r, "error.html", struct{ Error string }{strconv.Itoa(status) + err})
}
