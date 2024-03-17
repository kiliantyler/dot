package utils

// trimStringToLengthUTF8 adjusts a string to make it exactly the specified length in a UTF-8 safe manner.
// If the string is longer, it will be trimmed; if shorter, it will be padded with the specified padChar.
// Assumes padChar is a valid single-rune string.
// * We can't log here since this is called by the logging function, it would be a forever loop *
func TrimStringToLengthUTF8(s string, targetLength int) string {
	// Convert the string to a slice of runes to handle multibyte characters correctly.
	runes := []rune(s)
	padRunes := []rune(" ")

	if len(runes) > targetLength {
		return string(runes[:targetLength])
	}

	// Pad the string with the padChar rune if it's shorter than targetLength.
	for len(runes) < targetLength {
		runes = append(runes, padRunes[0])
	}

	return string(runes)
}
