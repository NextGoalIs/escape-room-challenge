package system

import (
	"bufio"
	"os"
	"strings"
)

func Scan() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
