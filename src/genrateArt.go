package rev

var (
	nbrSpaces int
	artLen    int
	asciiArt  string
)

func GenrateArt(asciiMap map[rune][]string, splittedInput []string) string {
	for _, line := range splittedInput {
		if line == "" {
			asciiArt += "\n"
			continue
		}
		asciiSlice := [][]string{}
		for index, char := range line {
			if char >= 32 && char <= 126 {
				if alignOption == "justify" && char == 32 {
					space := []string{"+", "+", "+", "+", "+", "+", "+", "+"}
					asciiSlice = append(asciiSlice, space)
					nbrSpaces++
				} else if oColor != "" {
					isColored := color(asciiMap[char], line, index)
					asciiSlice = append(asciiSlice, isColored)
				} else {
					asciiSlice = append(asciiSlice, asciiMap[char])
					artLen += len(asciiMap[char][0])
				}
			}
		}

		for i := 0; i < 8; i++ {
			if alignOption == "center" || alignOption == "reight" {
				align(alignOption)
			}
			for _, Art := range asciiSlice {
				if alignOption == "justify" && Art[i] == "+" {
					align(alignOption)
				} else {
					asciiArt += Art[i]
				}
			}
			asciiArt += "\n"
		}

		nbrSpaces = 0
		artLen = 0
	}
	return asciiArt
}
