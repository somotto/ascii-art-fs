package main

import (
	"fmt"
	"os"

	"asciifunc/ascii"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	words := os.Args[1]
	var banner string
	filenm := ""

	if len(os.Args) == 3 {
		banner = os.Args[2]
	} else {
		filenm = "standard.txt"
	}

	switch banner {
	case "shadow":
		filenm = "shadow.txt"

	case "thinkertoy":
		filenm = "thinkertoy.txt"
	default:
		filenm = "standard.txt"
	}

	words = ascii.ReplaceNonPrintChar(words)
	if ascii.ContainNonPrintChar(words) {
		fmt.Println("contains non-printable characters")
		return
	}
	if ascii.SpecialCases(words) {
		return
	}
	ascii.ProcessWords(words, filenm)
}
