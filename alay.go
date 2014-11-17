package alay

import (
	"strings"
	"unicode"
)

var m = map[rune]rune{
	'a': '4',
	'A': '4',
	'i': '1',
	'I': '1',
	'e': '3',
	'E': '3',
	'o': '0',
	'O': '0',
}

var mr map[rune]rune

func toAlayMapper(r rune) rune {
	if v, ok := m[r]; ok {
		return v
	}
	return r
}

func ToAlay(s string) string {
	return strings.Map(toAlayMapper, s)
}

func reverseM() {
	if mr == nil {
		mr = make(map[rune]rune)
		for k, v := range m {
			mr[v] = unicode.ToLower(k)
		}
	}
}

func normalizeMapper(r rune) rune {
	reverseM()
	if v, ok := mr[r]; ok {
		return v
	}
	return r
}

func Normalize(s string) string {
	return strings.Map(normalizeMapper, s)
}
