package main

import (
	"github.com/anatolethien/ascii-art"
	"github.com/anatolethien/go-reloaded"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var t, _ = template.ParseFiles("./templates/index.html")
		t.Execute(w, nil)
	})
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		var input = reloaded.Split(r.FormValue("input"), "\\n")
		var banner = r.FormValue("banner")
		type display struct{ Output string }
		var ascii = display{Output: ascii.Ascii(input, banner)}
		var t, _ = template.ParseFiles("./templates/index.html")
		t.Execute(w, ascii)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":3000", nil)
}
