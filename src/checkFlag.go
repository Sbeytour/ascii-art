package rev

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func checkFlag(arg string) {
	flag = false
	colorRgb = false
	if strings.HasPrefix(arg, "--output=") {
		flag = true
		OutputFile = arg[9:]
		if !strings.HasSuffix(arg, ".txt") || OutputFile == ".txt" {
			log.Fatalln("Error: OUTPUT file name incorrect")
		}
	}

	if strings.HasPrefix(arg, "--align=") {
		flag = true
		alignOption = arg[8:]
		option := []string{"center", "right", "left", "justify"}
		if !contain(alignOption, option) {
			log.Fatalln("Error: align option incorrect")
		}
	}

	if strings.HasPrefix(arg, "--color=") {
		flag = true
		oColor = arg[8:]
		rgxp := regexp.MustCompile(`^rgb\((\d+),\s?(\d+),\s?(\d+)\)$`)
		if rgxp.MatchString(oColor) {
			colorRgb = true
			submatched := rgxp.FindStringSubmatch(oColor)
			for i := 1; i < len(submatched); i++ {
				nbr, err := strconv.Atoi(submatched[i])
				if err != nil {
					log.Fatalln("Error in converting")
				}
				if nbr > 255 || nbr < 0 {
					log.Fatalln("error rgb number is incorrect")
				}
			}
			R = submatched[1]
			G = submatched[2]
			B = submatched[3]
		} else {
			colorOption := []string{"red", "green", "yallow", "blue", "orange", "pink", "purple", "aqua"}
			if !contain(oColor, colorOption) {
				log.Fatalln("Error: color doesn't exist")
			}
		}

	}

	if strings.HasPrefix(arg, "--reverse=") {
		ReverseFile = arg[10:]
		if !strings.HasSuffix(ReverseFile, ".txt") || ReverseFile == ".txt" {
			log.Fatalln("Error: REVERSE file name incorrect")
		}
	}
}

func contain(str string, slice []string) bool {
	for _, typ := range slice {
		if typ == str {
			return true
		}
	}
	return false
}
