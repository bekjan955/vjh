package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

func Convert(input string) string {
	if isMorse(input) {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}

func isMorse(input string) bool {
	for _, r := range input {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			return false
		}
	}
	return true
}
