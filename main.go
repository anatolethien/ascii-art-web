package main

import (
	"./pkg/ascii_art_web"
	"fmt"
)

func main() {

	ascii_art_web.OpenBrowser(ascii_art_web.Url)
	fmt.Printf("Server loading on \"%s\"\n", ascii_art_web.Url)
	ascii_art_web.WebServer()

}
