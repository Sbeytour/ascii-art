package rev

import (
	"os"
	"strings"
)

func ManageRevFile(file string) ([]string, error) {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	fileSplitted := strings.Split(string(fileContent), "\n")

	return fileSplitted, nil
}
