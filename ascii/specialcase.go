package ascii

import "fmt"

// Checks and enables the running of
// newline and empty string.
func SpecialCases(words string) bool {
	if words == "\\n" {
		fmt.Println()
		return true
	} else if words == "" {
		return true
	}
	return false
}
