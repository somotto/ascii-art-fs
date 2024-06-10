package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintChar(word, filename string) string {
	out := ""
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading from file")
		os.Exit(1)
	}

	filecontent := strings.Split(string(file), "\n")

	if strings.Contains(string(file), "o") {
		filecontent = strings.Split(string(file), "\r\n")
	}

	for l := 1; l <= 8; l++ {
		for _, char := range word {
			index := int(char-32) * 9
			out += filecontent[index+l]
		}
		out += "\n"
	}

	return out
}
