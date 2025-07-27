package functions

import (
	"log"
	"os"
	"strings"
)

// a function to read the input and store the ascii values in a map
func AsciiArtTable(banner string) map[int]string {
	bannerName := "banners/" + banner + ".txt"
	fileData, err := os.ReadFile(bannerName)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	blocks := strings.Split(string(fileData), "\n\n")

	if strings.Count(blocks[0], "\n") == 8 {
		blocks[0] = blocks[0][1:]
	}
	asciiArtTable := make(map[int]string)

	key := 32
	for _, block := range blocks {
		asciiArtTable[key] = block
		key++
	}
	return asciiArtTable
}
