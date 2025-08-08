// Package helpers provides utility functions for HTTP request handling,
// template rendering, and error page management in Go web applications.
package helpers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var Help Helpers

// Configuration variables
var (
	PagesPath = "helpers/pages/"
)

// Data represents the structure for error page template data.
type Data struct {
	Error   string // Human-readable error message (e.g., "Not Found")
	StatusE string // HTTP status code as string (e.g., "404")
}

// Helpers is an empty struct that serves as a receiver for helper methods.
type Helpers struct{}

// CheckGet validates that the incoming HTTP request uses the GET method.
// If the request method is not GET, it automatically renders a 405 Method Not Allowed
// error page and terminates the request processing.
func (a Helpers) CheckGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		a.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
}

// ErrorPage renders a custom error page with the specified HTTP status code.
func (a Helpers) ErrorPage(w http.ResponseWriter, status int) {
	msg := http.StatusText(status)

	tmp, err := template.ParseFiles(PagesPath + "error.html")
	if err != nil {
		http.Error(w, strconv.Itoa(status)+" "+msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Use a buffer to execute the template safely before writing to response
	var buf bytes.Buffer
	errExec := tmp.Execute(&buf, Data{Error: msg, StatusE: strconv.Itoa(status)})
	if errExec != nil {
		a.InternalServerError(w)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// RenderPage renders any HTML template page with provided data.
// This is a generic template rendering method that can be used for any page
// Note: The template file path is constructed by appending ".html" to UrlPage,
func (a Helpers) RenderPage(UrlPage string, RendredData any, w http.ResponseWriter) {
	// Parse and compile the template - panics if parsing fails
	Templ := template.Must(template.ParseFiles(UrlPage + ".html"))
	err := Templ.Execute(w, RendredData)
	if err != nil {
		a.ErrorPage(w, http.StatusInternalServerError)
		return
	}
}

// sends HTTP GET request to URL,
// decodes the JSON response into 'result',
// and returns any error if exist.
func (a Helpers) Fetch(url string, result interface{}) error {
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

// InternalServerError renders 500 Internal Server Error page with embedded HTML and CSS.
func (Helpers) InternalServerError(w http.ResponseWriter) {
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

