package rev

import "strings"

func SplitInput(input string) []string {
	splittedInput := strings.Split(input, `\n`)
	if newLines(splittedInput) {
		splittedInput = splittedInput[:len(splittedInput)-1]
	}

	return splittedInput
}

func newLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return false
		}
	}
	return true
}
