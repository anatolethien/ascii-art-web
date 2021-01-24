package ascii_art_func

import (
	"bufio"
	"fmt"
	"os"
)

func Mysterious(text string) {

	if text == "Lizardon" {
		var file, _ = os.Open("assets/charizard.txt")
		var scanner = bufio.NewScanner(file)
		for scanner.Scan() == true {
			fmt.Println(scanner.Text())
		}
	}

}
