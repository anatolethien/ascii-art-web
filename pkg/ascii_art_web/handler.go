package ascii_art_web

import (
	"../ascii_art_err"
	"../ascii_art_func"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		w.WriteHeader(ascii_art_err.OK.Status)
		var t, _ = template.ParseFiles("templates/index.html")
		t.Execute(w, nil)

	} else if r.Method == "POST" {

		var text = r.FormValue("text")
		var banner = r.FormValue("banner")

		if ascii_art_func.ValidText(text) == true {

			var display = Display{
				Input:  text,
				Output: ascii_art_func.Fs(text, banner),
			}

			w.WriteHeader(ascii_art_err.Created.Status)
			var t, _ = template.ParseFiles("templates/index.html")
			t.Execute(w, display)

		} else if ascii_art_func.ValidText(text) == false {

			w.WriteHeader(ascii_art_err.BadRequest.Status)
			var t, _ = template.ParseFiles("templates/error.html")
			t.Execute(w, ascii_art_err.BadRequest)

		}

	}

}
