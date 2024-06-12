package ascii

import (
	"fmt"
	"strings"
)

// Enables the processing of the string passed to
// give the output in representation of the banner file stated.
func ProcessWords(words string, filename string) {
	wordsArr := strings.Split(words, "\\n")

	for _, word := range wordsArr {
		if word == "" {
			fmt.Println()
			continue
		}
		fmt.Print(PrintChar(word, filename))
	}
}
