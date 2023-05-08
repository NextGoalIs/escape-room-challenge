package utils

import "strings"

func GetStringCenter(s string, PADDING int) string {
	if len(s) >= PADDING {
		return s
	}
	return strings.Repeat(" ", (PADDING-len(s))/2) + s
}