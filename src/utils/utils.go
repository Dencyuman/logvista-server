package utils

import (
	strings "strings"
)

func ReplaceDummyLf(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}
