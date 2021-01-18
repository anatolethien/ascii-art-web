package ascii_art_web

import (
	"html/template"
	"net/http"
)

func WebServer()  {

	t = template.Must(template.ParseGlob("www/index.html"))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)

}
