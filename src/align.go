package rev

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func align(option string) {
	terWidth := TerminalSize()
	switch option {
	case "reight":
		nbr := terWidth - artLen
		asciiArt += strings.Repeat(" ", nbr)
	case "center":
		nbr := (terWidth - artLen) / 2
		asciiArt += strings.Repeat(" ", nbr)
	case "justify":
		nbr := (terWidth - artLen) / nbrSpaces
		asciiArt += strings.Repeat(" ", nbr)
	}
}

func TerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln("error of executing the command")
	}

	tsize := strings.Fields(string(out))

	size, _ := strconv.Atoi(tsize[1])
	return size
}
