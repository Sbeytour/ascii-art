package rev

import (
	"fmt"
	"log"
	"os"
)

var (
	flag        bool
	OutputFile  string
	bannerFile  string
	input       string
	alignOption string
	oColor      string
	subStr      string
	ReverseFile string
)

func CheckArgs(args []string) (string, string) {
	if len(args) < 2 || len(args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

	checkFlag(args[1])

	if (len(args) == 2 && flag) || (len(args) == 4 && !flag) || (len(args) == 5 && oColor == "") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

	if len(os.Args) == 5 {
		bannerFile = banners(args[4])
		subStr = args[2]
	} else if len(args) == 4 && oColor != "" {
		if args[3] == "standard" || args[3] == "thinkertoy" || args[3] == "shadow" {
			bannerFile = banners(args[3])
		} else {
			bannerFile = "banners/standard.txt"
			subStr = args[2]
		}
	} else if len(args) == 4 {
		bannerFile = banners(args[3])
	} else if len(args) == 3 && !flag {
		bannerFile = banners(args[2])
	} else {
		bannerFile = "banners/standard.txt"
	}

	if subStr != "" {
		input = args[3]
	} else if flag {
		input = args[2]
	} else {
		input = args[1]
	}
	return bannerFile, input
}

func banners(arg string) string {
	switch arg {
	case "standard":
		bannerFile = "banners/standard.txt"
	case "shadow":
		bannerFile = "banners/shadow.txt"
	case "thinkertoy":
		bannerFile = "banners/thinkertoy.txt"
	default:
		log.Fatalln("invalid banner")
	}
	return bannerFile
}
