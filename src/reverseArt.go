package rev

var rev string

func Reverse(inputFile []string, asciiMap map[rune][]string) string {
	for i := 0; i < len(inputFile); i++ {
		if (inputFile[i] == "" || inputFile[i] == "$") && i != len(inputFile)-1 {
			rev += "\n"
			continue
		}

		if len(inputFile[i:]) >= 8 {
			for char := ' '; char <= '~'; char++ {
				charWidth := len(asciiMap[char][0])
				if isMatched(asciiMap[char], inputFile[i:], charWidth) {
					rev += string(char)
					inputFile = removeChar(charWidth, inputFile, i)
					char = 31
				}
			}

            if isNotValidFormat(inputFile[i:]) {
                return "file format incorrect"
            }

			i += 7
			rev += "\n"
		}
	}
	return rev[:len(rev)-1]
}

func isNotValidFormat(inputFile []string) bool {
	for i := 0; i < 8; i++ {
		if inputFile[i] != "" && inputFile[i] != "$" {
			return true
		}
	}
	return false
}

func isMatched(asciiChar, inputFile []string, charWidth int) bool {
	for i := 0; i < 8; i++ {
		if len(inputFile[i]) < len(asciiChar[i]) || asciiChar[i] != inputFile[i][:charWidth] {
			return false
		}
	}
	return true
}

func removeChar(width int, inputFile []string, idx int) []string {
	newInputFile := []string{}
	newInputFile = append(newInputFile, inputFile[:idx]...)

	for i := 0; i < 8; i++ {
		newInputFile = append(newInputFile, inputFile[idx+i][width:])
	}
	newInputFile = append(newInputFile, inputFile[idx+8:]...)
	return newInputFile
}
