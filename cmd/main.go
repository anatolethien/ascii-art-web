package main

import (
	"fmt"
	"github.com/anatolethien/ascii-art"
	"github.com/anatolethien/go-reloaded"
	"html/template"
	"net/http"
)

type display struct {
	Output string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var ascii = display{Output: ascii.Ascii(reloaded.Split(" ", "\\n"), "standard")}
		var t, _ = template.ParseFiles("./templates/index.html")
		t.Execute(w, ascii)
	})
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		var input, banner = reloaded.Split(r.FormValue("input"), "\\n"), r.FormValue("banner")
		var ascii = display{Output: ascii.Ascii(input, banner)}
		var t, _ = template.ParseFiles("./templates/index.html")
		t.Execute(w, ascii)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Printf("Launching web app at http://127.0.0.1:3000...\n")
	http.ListenAndServe(":3000", nil)
}
