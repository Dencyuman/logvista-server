package utils

import (
	strings "strings"
)

// 文字列中のダミー改行文字\\nを改行文字に置き換える
func ReplaceDummyLf(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}

// string配列から空文字を除外する
func FilterEmptyStrings(strings []string) []string {
	var result []string
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
