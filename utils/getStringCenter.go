package utils

import "strings"

func GetStringCenter(s string, INTERVAL int) string {
	if len(s) >= INTERVAL {
		return s
	}
	return strings.Repeat(" ", (INTERVAL-len(s))/2) + s
}