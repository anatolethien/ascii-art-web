package ascii_art_func

import (
	"bufio"
	"fmt"
	"os"
)

func Mysterious(text string) {

	if text == "lizardon" {
		var file, _ = os.Open("assets/mysterious/charizard.txt")
		var scanner = bufio.NewScanner(file)
		for scanner.Scan() == true {
			fmt.Println(scanner.Text())
		}
	} else if text == "doom" {
		var file, _ = os.Open("assets/mysterious/doom.txt")
		var scanner = bufio.NewScanner(file)
		for scanner.Scan() == true {
			fmt.Println(scanner.Text())
		}
	}

}
