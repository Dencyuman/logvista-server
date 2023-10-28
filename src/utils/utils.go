package utils

import (
	strings "strings"
)

func PointerToInt(n int) *int {
	return &n
}

func ReplaceDummyLf(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}
