package ascii

import (
	"fmt"
	"os"
	"strings"
)

// Reads the file and generates the ascii-art
// the way it is from the banner files.
func PrintChar(word, filename string) string {
	out := ""
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading from file")
		os.Exit(1)
	}
	// Handling empty files

	if len(strings.Split(string(file), "\n")) != 856 {
		fmt.Println("Error: >> Banner filecontent  is incomplete with length of: ", len(strings.Split(string(file), "\n")))
		os.Exit(0)
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
