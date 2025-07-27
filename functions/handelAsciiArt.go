package functions

// handling the validity of the input
func HandelAsciiArt(inputText, banner string) (string, bool) {
	if !CheckRange(inputText) {
		return "", true
	}

	if !CheckBanner(banner) {
		return "", true
	}

	wordsSlice := SplitInput(inputText)

	return AppendArt(wordsSlice, AsciiArtTable(banner)), false
}
