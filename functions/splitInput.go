package functions

import (
	"strings"
)

// ParseInput splits the input string by "\n" and removes a leading empty string if present.
func SplitInput(input string) []string {
	wordsSlice := strings.Split(input, "\r\n")
	for i, val := range wordsSlice {
		if val != "" {
			break
		}
		if i == len(wordsSlice)-1 && val == "" {
			wordsSlice = wordsSlice[1:]
		}
	}
	return strings.Split(input, "\r\n")
}
