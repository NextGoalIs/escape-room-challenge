package system

import (
	"bufio"
	"strings"
)

func Scan(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
