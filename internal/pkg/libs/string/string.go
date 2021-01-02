package _stringUtils

import (
	"strings"
	"unicode"
)

func TrimAll(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.TrimSpace(str)

	return str
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func FindInArr(str string, arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}

	return false
}
