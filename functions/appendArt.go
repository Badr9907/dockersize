package functions

import (
	"strings"
)

// a function to range in the input and print it if matches with the map values.
func AppendArt(wordsSlice []string, asciiArtTable map[int]string) string {
	artSlice := [][]string{}
	asciiArtResult := ""
	for _, word := range wordsSlice {
		if len(word) == 0 {
			asciiArtResult += "\n"
			continue
		}

		for _, letter := range word {
			for key, value := range asciiArtTable {
				if letter == rune(key) {
					artSlice = append(artSlice, strings.Split(value, "\n"))
				}
			}
		}

		for i := 0; i < 8; i++ {
			lineParts := []string{}
			for _, line := range artSlice {
				lineParts = append(lineParts, line[i])
			}
			asciiArtResult = asciiArtResult + strings.Join(lineParts, "") + "\n"
		}
		artSlice = nil
	}
	return asciiArtResult
}
