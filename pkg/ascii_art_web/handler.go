package ascii_art_web

import (
	"../ascii_art_err"
	"../ascii_art_func"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		w.WriteHeader(ascii_art_err.OK.Status)
		var t, _ = template.ParseFiles("templates/index.html")
		t.Execute(w, nil)

	} else if r.Method == "POST" {

		var text = r.FormValue("text")
		var banner = r.FormValue("banner")
		var mode = r.FormValue("mode")

		if ascii_art_func.ValidText(text) == true {

			display = Display{
				Text:   text,
				Result: ascii_art_func.Fs(text, banner),
			}

			if mode == "show" {

				w.WriteHeader(ascii_art_err.Created.Status)
				var t, _ = template.ParseFiles("templates/index.html")
				t.Execute(w, display)

			} else if mode == "download" {

				var output = strings.NewReader(display.Result)
				w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
				w.Header().Set("Content-Length", strconv.Itoa(len(display.Result)))
				io.Copy(w, output)

			}

		} else if ascii_art_func.ValidText(text) == false {

			w.WriteHeader(ascii_art_err.BadRequest.Status)
			var t, _ = template.ParseFiles("templates/error.html")
			t.Execute(w, ascii_art_err.BadRequest)

		}

	}

}
