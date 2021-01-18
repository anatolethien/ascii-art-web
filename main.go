package main

import (
	"./pkg/ascii_art_web"
	"fmt"
)

func main() {

	fmt.Println("localhost:1337")
	ascii_art_web.WebServer()

}
