package utils

import (
	"os"
	"os/exec"
)

func ClearConsoleLinux() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
