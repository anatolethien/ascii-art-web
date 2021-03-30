package ascii_art_web

import (
	"net/http"
)

func WebServer() {

	http.HandleFunc("/", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("templates/media"))))
	http.ListenAndServe(":80", nil)

}
