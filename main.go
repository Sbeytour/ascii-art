package main

import (
	"fmt"
	"log"
	"os"

	"reverse/src"
)

func main() {
	args := os.Args
	bannerFile, input := rev.CheckArgs(args)
	if rev.ReverseFile != "" {
		asciiMap, err := rev.CreateMap("banners/standard.txt")
		if err != nil {
			log.Fatalln("Error of reading the file and crating Map")
		}

		inputFile, err := rev.ManageRevFile(rev.ReverseFile)
		if err != nil {
			log.Fatalln("Error of reading the input file")
		}
	
		printing := rev.Reverse(inputFile, asciiMap)
		fmt.Println(printing)
	} else {
		asciiArt := ""
		asciiMap, err := rev.CreateMap(bannerFile)
		if err != nil {
			log.Fatalln("Error of reading the file")
		}

		if input == "" {
			os.Remove(rev.OutputFile)
			os.Exit(0)
		}

		splittedInput := rev.SplitInput(input)
		asciiArt = rev.GenrateArt(asciiMap, splittedInput)
		err = rev.CreateFile(asciiArt)
		if err != nil {
			log.Fatalln("error of creating file")
		}

	}
}
