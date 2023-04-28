package utils

import "strings"

const INTERVAL = 80

func PrintStringCenter(s string) string {
	if len(s) >= INTERVAL {
		return s
	}
	return strings.Repeat(" ", (INTERVAL-len(s))/2) + s
}