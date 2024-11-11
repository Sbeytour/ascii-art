package rev

import (
	"fmt"
	"os"
)

func CreateFile(asciiart string) error {
	if OutputFile != "" {
		file, err := os.Create(OutputFile)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(asciiart)
		if err != nil {
			return err
		}
	} else {
		fmt.Print(asciiart)
	}
	return nil
}
