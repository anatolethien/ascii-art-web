package ascii_art_func

import "strings"

func Fs(text string, banner string) string {

	var output []byte
	var index int

	var file = scanner(banner)

	var splitText []string
	splitText = strings.Split(text, "\n")

	for _, word := range splitText {
		output = generator(output, file, word, index)
	}

	return string(output)

}
