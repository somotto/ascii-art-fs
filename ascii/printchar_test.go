package ascii

import (
	"strings"
	"testing"
)

func TestPrintChar(t *testing.T) {
	expected := []string{
		" _    _  ",
		"| |  | | ",
		"| |__| | ",
		"|  __  | ",
		"| |  | | ",
		"|_|  |_| ",
		"         ",
		"         ",
	}
	input := "../standard.txt"
	output := PrintChar("H", input)
	exp := strings.Join(expected, "\n") + "\n"
	if output != exp {
		t.Fatalf("got: \n%v\n expected: \n%v\n", output, exp)
	}
}

func TestContainNonPrintChar(t *testing.T) {
	input := "你好"
	expected := true
	result := ContainNonPrintChar(input)
	if result != expected {
		t.Fatalf("got: \n%v\n want: \n%v\n", result, expected)
	}
}

func TestSpecialCases(t *testing.T) {
	input := "\\n"
	expected := true
	result := SpecialCases(input)
	if result != expected {
		t.Fatalf("got: \n%v\n want: \n%v\n", result, expected)
	}
}

func TestReplaceNonPrintChar(t *testing.T) {
	input := "\\t"
	expected := "    "
	result := ReplaceNonPrintChar(input)
	if result != expected {
		t.Fatalf("got: \n%v\n want: \n%v\n", result, expected)
	}
}
