package main

import (
	"fmt"
	"github.com/anatolethien/ascii-art"
	"github.com/anatolethien/go-reloaded"
	"html/template"
	"net/http"
)

const Port = ":3000"

type display struct {
	Output string
}

type error struct {
	Status int
	Error  string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, error{Status: http.StatusNotFound, Error: http.StatusText(http.StatusNotFound)})
		return
	}
	var input, banner = reloaded.Split(" ", "\\n"), "standard"
	var ascii = display{Output: ascii.Ascii(input, banner)}
	var t, _ = template.ParseFiles("./templates/index.html")
	t.Execute(w, ascii)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	var input, banner = reloaded.Split(r.FormValue("input"), "\\n"), r.FormValue("banner")
	var ascii = display{Output: ascii.Ascii(input, banner)}
	var t, _ = template.ParseFiles("./templates/index.html")
	t.Execute(w, ascii)
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(err.Status)
	var t, _ = template.ParseFiles("./templates/error.html")
	t.Execute(w, err)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Printf("Launching web app at http://127.0.0.1%s...\n", Port)
	http.ListenAndServe(Port, nil)
}
