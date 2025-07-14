package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if isMorse(data) {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}

func isMorse(s string) bool {

	valid := !strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
	if !valid {
		return false
	}

	return strings.ContainsAny(s, ".-")
}
