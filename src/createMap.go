package rev

import (
	"os"
	"strings"
)

func CreateMap(bannerfile string) (map[rune][]string, error) {
	fileContent, err := os.ReadFile(bannerfile)
	if err != nil {
		return nil, err
	}

	fileContent = cleaner(fileContent)
	splitFile := strings.Split(string(fileContent), "\n")

	asciiMap := make(map[rune][]string)
	dec := 31
	for _, line := range splitFile {
		if line == "" {
			dec++
		} else {
			asciiMap[rune(dec)] = append(asciiMap[rune(dec)], line)
		}
	}
	return asciiMap, nil
}

func cleaner(file []byte) []byte {
	res := []byte{}
	for _, byt := range file {
		if byt != '\r' {
			res = append(res, byt)
		}
	}
	return res
}
