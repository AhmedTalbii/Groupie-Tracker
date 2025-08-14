// Package helpers provides utility functions for HTTP request handling,
// template rendering, and error page management in Go web applications.
package biblio

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Helpers struct{} // Helpers is an empty struct that serves as a receiver for helper methods.

var Help Helpers

var PagesPath = "biblio/pages/" // PagesPath is the directory path for external files.

// Data represents the structure for error page template data.
type Data struct {
	Error   string
	StatusE string
}

// CheckGet ensures the request method is GET;
// if not, it sends a 405 error page and stops processing.
func (a *Helpers) CheckGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		a.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
}

// serves static files if the path exists and is not a directory;
// otherwise, it shows 404 or 403 error pages.
func (a *Helpers) StaticsHandler(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		a.ErrorPage(w, http.StatusNotFound)
		return
	}
	if info.IsDir() {
		a.ErrorPage(w, http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

// ErrorPage renders error page with the specified HTTP status code.
func (a *Helpers) ErrorPage(w http.ResponseWriter, status int) {
	msg := http.StatusText(status)

	if status == 404 {
		msg = "Page " + msg
	}

	tmp, err := template.ParseFiles(PagesPath + "error.html")
	if err != nil {
		a.InternalServerError(w)
		return
	}

	var buf bytes.Buffer // Use a buffer to execute the template safely before writing to response
	errExec := tmp.Execute(&buf, Data{Error: msg, StatusE: strconv.Itoa(status)})
	if errExec != nil {
		a.InternalServerError(w)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// renders an HTML page with data.
// Shows error page if rendering fails.
func (a *Helpers) RenderPage(UrlPage string, RendredData any, w http.ResponseWriter) {
	var buf bytes.Buffer
	if err := a.RenderPageInsideBuffer(UrlPage, RendredData, &buf); err != nil {
		a.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

// renders a template with data into a buffer.
// return err if parsing or execution fails.
func (a *Helpers) RenderPageInsideBuffer(UrlPage string, RendredData any, buf *bytes.Buffer) error {
	Templ, err := template.ParseFiles(UrlPage + ".html")
	if err != nil {
		return err
	}

	err = Templ.Execute(buf, RendredData)
	if err != nil {
		return err
	}
	return nil
}

// sends a GET request to URL, decodes JSON response into result, and returns any error.
func (a *Helpers) Fetch(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

// renders 500 Internal Server Error page with embedded HTML and CSS.
func (a *Helpers) InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title></title>
		</head>

		 <style>
		        html,
		        body {
		            color: white;
		            font-family: Arial, Helvetica, sans-serif;
		            background-color: rgb(48, 47, 47);
		            height: 100vh;
		            display: flex;
		            justify-content: center;
		            align-items: center;
		            flex-direction: column;
		        }



		        .err {
		            color: red;
		            font-size: 40px;
		            font-weight: bold;
		            margin-bottom: 20px;
		        }
		    </style>
		<body>
		    <div class="err">500 Status Internal Server Error</div>
		</body>
		</html>`))
}

// check the connection if is on
// else return error
func (a *Helpers) CheckConnection() error {
	_, err := http.Get("https://google.com")
	if err != nil {
		return err
	}
	return nil
}
