package ascii_art_func

func ValidText(text string) bool {

	for _, v := range text {

		if v <= 126 {
			continue
		} else {
			return false
		}

	}

return true
}
