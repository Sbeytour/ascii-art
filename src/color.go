package rev

import (
	"fmt"
)

var colors = map[string]string{
	"red":     "\033[38;2;255;0;0m",
	"green":   "\033[38;2;0;255;0m",
	"yallow":  "\033[38;2;255;255;0m",
	"blue":    "\033[38;2;0;0;255m",
	"orange":  "\033[38;2;255;127;0m",
	"Fuchsia": "\033[38;2;255;0;255m",
	"purple":  "\033[38;2;102;0;204m",
	"aqua":    "\033[38;2;0;255;255m",
}

var (
	R, G, B  string
	colorRgb bool
)

func color(art []string, line1 string, index int) []string {
	resetCode := "\033[0m"
	coloring := colors[oColor]
	coloredArt := []string{}
	indice := []int{}
	if subStr != "" {
		for i := 0; i <= len(line1)-len(subStr); i++ {
			if subStr == line1[i:i+len(subStr)] {
				for j := i; j < i+len(subStr); j++ {
					indice = append(indice, j)
				}
			}
		}
	}

	for _, line := range art {
		if subStr == "" {
			if colorRgb {
				colored1 := fmt.Sprintf("\033[38;2;%s;%s;%sm", R, G, B) + line + resetCode
				coloredArt = append(coloredArt, colored1)
			} else {
				colored2 := coloring + line + resetCode
				coloredArt = append(coloredArt, colored2)
			}
		} else {
			if matching(index, indice) {
				if colorRgb {
					colored1 := fmt.Sprintf("\033[38;2;%s;%s;%sm", R, G, B) + line + resetCode
					coloredArt = append(coloredArt, colored1)
				} else {
					colored2 := coloring + line + resetCode
					coloredArt = append(coloredArt, colored2)
				}
			} else {
				coloredArt = append(coloredArt, line)
			}
		}
	}

	return coloredArt
}

func matching(i int, ind []int) bool {
	for _, j := range ind {
		if j == i {
			return true
		}
	}
	return false
}
