package helper

import "strings"

func RemoveChars(input string, charsToRemove string) string {
	return strings.Map(func(r rune) rune {
		// Check if the rune is in charsToRemove
		for _, char := range charsToRemove {
			if r == char {
				return -1 // Remove the character
			}
		}
		return r
	}, input)
}
