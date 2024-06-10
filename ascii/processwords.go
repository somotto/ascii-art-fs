package ascii

import (
	"fmt"
	"strings"
)

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
