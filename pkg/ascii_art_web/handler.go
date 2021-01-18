package ascii_art_web

import (
	"../ascii_art_func"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		if r.URL.Path != "/" {
			return
		}

		t.ExecuteTemplate(w, "index.html", nil)

	} else if r.Method == "POST" {

		var text = r.FormValue("text")
		var banner = r.FormValue("banner")

		var result = ascii_art_func.Fs(text, banner)

		var display = Display{
			Input: text,
			Output: result,
		}

		t.ExecuteTemplate(w, "index.html", display)

	}

}
