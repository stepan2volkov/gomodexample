package gomodexample

import (
	"unicode/utf8"
)

func ReverseStringFast(value string) string {
	chars := []rune(value)
	stringLength := utf8.RuneCountInString(value)

	for i := 0; i < stringLength/2; i++ {
		chars[i], chars[stringLength-i-1] = chars[stringLength-i-1], chars[i]
	}
	return string(chars)
}
