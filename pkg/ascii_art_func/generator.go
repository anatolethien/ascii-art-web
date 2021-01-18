package ascii_art_func

func generator(output []byte, file []byte, word string, index int) []byte {

	if index == 8 {
		return output
	}

	for i, n := 0, len(word); i < n; i++ {

		lineNum := 0
		a := int(word[i]-32)*9 + 2 + index

		for j, l := 0, len(file); j < l; j++ {

			if file[j] == 13 {
				continue
			}

			if file[j] == 10 {
				lineNum++
			} else if lineNum == a-1 {
				output = append(output, file[j])
			} else if lineNum == a {
				break
			}

		}

	}

	output = append(output, '\n')

	return generator(output, file, word, index+1)
}
