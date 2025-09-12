package main

import (
	"fmt"
	"regexp"
	"strings"
)

func removeComments(source []string) []string {
	var output []string

	code := strings.Join(source, "\n")

	block := regexp.MustCompile(`(?s)/\*.*?\*/`)
	code = block.ReplaceAllString(code, "")

	block = regexp.MustCompile(`//.*`)
	code = block.ReplaceAllString(code, "")

	for _, line := range strings.Split(code, "\n") {
		if strings.TrimSpace(line) != "" {
			output = append(output, line)
		}
	}

	return output
}

func FindReplace() {
	source := []string{"/* Example code for feature */", "int main() {", "  /*", "  This is a", "  block comment", "  */", "  int value = 10;  // This is an inline comment", "  int sum = value + /* this is // also a block */ value;", "  return 0;", "}"}
	for _, line := range removeComments(source) {
		fmt.Printf("%#v\n", line)
	}
}
