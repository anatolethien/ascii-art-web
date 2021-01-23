package ascii_art_web

var display Display

type Display struct {
	Text   string
	Result string
}

type Status struct {
	Status     int
	StatusText string
}

var Url = "http://localhost:1337"
