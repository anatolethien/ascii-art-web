package ascii_art_web

import "html/template"

var t *template.Template

type Display struct {
	Input string
	Output string
}

var Url = "http://localhost:1337"
