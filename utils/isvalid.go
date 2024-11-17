package utils

import (
	"regexp"
	"strings"
)

func Isvalid(st string) bool {
	st = strings.ToLower(st)
	if len(st) != 6 || !isAlphaNumeric(st) || strings.Contains(st, "https") || strings.Contains(st, "http") {
		return false
	}
	return true
}

func isAlphaNumeric(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}
