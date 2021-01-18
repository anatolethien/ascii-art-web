package ascii_art_func

import (
	"io/ioutil"
)

func scanner(banner string) []byte {

	var file, err = ioutil.ReadFile("./assets/banner/" + banner + ".txt")
	if err != nil {
		return []byte("")
	}

	return file
}
